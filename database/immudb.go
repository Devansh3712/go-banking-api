package database

import (
	"context"
	"crypto/rand"
	"fmt"
	"strconv"
	"time"

	"github.com/Devansh3712/go-banking-api/config"
	"github.com/Devansh3712/go-banking-api/models"
	"github.com/codenotary/immudb/pkg/client"
	"google.golang.org/grpc/metadata"
)

// Create a random transaction ID.
func generateTxnID() (*string, error) {
	bytes := make([]byte, 8)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	id := fmt.Sprintf("%x", bytes)
	return &id, nil
}

// Create transaction table for storing withdrawal
// and deposits of a user.
func MigrateImmuDB() error {
	connection, err := client.NewImmuClient(client.DefaultOptions())
	if err != nil {
		return err
	}
	ctx := context.Background()
	response, err := connection.Login(
		ctx,
		[]byte(config.GetEnv("IMMUDB_USERNAME")),
		[]byte(config.GetEnv("IMMUDB_PASSWORD")),
	)
	if err != nil {
		return err
	}
	md := metadata.Pairs("authorization", response.Token)
	ctx = metadata.NewOutgoingContext(ctx, md)
	_, err = connection.SQLExec(
		ctx,
		`CREATE TABLE IF NOT EXISTS transactions (
			id 				VARCHAR[16],
			type 			VARCHAR,
			amount 			VARCHAR,
			account_number 	VARCHAR,
			time			TIMESTAMP,
			PRIMARY KEY (id)
		)`,
		map[string]interface{}{},
	)
	if err != nil {
		return err
	}
	return nil
}

func CreateTransaction(_type models.Txn, amount string, accNumber string) (*string, error) {
	connection, err := client.NewImmuClient(client.DefaultOptions())
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	response, err := connection.Login(
		ctx,
		[]byte(config.GetEnv("IMMUDB_USERNAME")),
		[]byte(config.GetEnv("IMMUDB_PASSWORD")),
	)
	if err != nil {
		return nil, err
	}
	md := metadata.Pairs("authorization", response.Token)
	ctx = metadata.NewOutgoingContext(ctx, md)
	txnID, err := generateTxnID()
	if err != nil {
		return nil, err
	}
	_, err = connection.SQLExec(
		ctx,
		`INSERT INTO transactions (id, type, amount, account_number, time)
		VALUES (@id, @type, @amount, @accNumber, @currTime)`,
		map[string]interface{}{
			"id":        *txnID,
			"type":      _type.Value(),
			"amount":    amount,
			"accNumber": accNumber,
			"currTime":  time.Now(),
		},
	)
	if err != nil {
		return nil, err
	}
	return txnID, nil
}

func GetTransactions(accNumber string, limit int) (*[]models.Transaction, error) {
	connection, err := client.NewImmuClient(client.DefaultOptions())
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	response, err := connection.Login(
		ctx,
		[]byte(config.GetEnv("IMMUDB_USERNAME")),
		[]byte(config.GetEnv("IMMUDB_PASSWORD")),
	)
	if err != nil {
		return nil, err
	}
	md := metadata.Pairs("authorization", response.Token)
	ctx = metadata.NewOutgoingContext(ctx, md)
	result, err := connection.SQLQuery(
		ctx,
		fmt.Sprintf(`
		SELECT * FROM transactions
		WHERE account_number = @accNumber
		LIMIT %d`, limit),
		map[string]interface{}{"accNumber": accNumber},
		true,
	)
	if err != nil {
		return nil, err
	}
	var transactions []models.Transaction
	for _, row := range result.Rows {
		amountFloat, _ := strconv.ParseFloat(row.Values[2].GetS(), 64)
		txn := models.Transaction{
			TxnID:     row.Values[0].GetS(),
			Type:      row.Values[1].GetS(),
			Amount:    amountFloat,
			Number:    row.Values[3].GetS(),
			Timestamp: time.UnixMicro(row.Values[4].GetTs()),
		}
		transactions = append(transactions, txn)
	}
	return &transactions, nil
}

// Get withdrawals or deposits of a user.
func GetTransactionsByType(_type models.Txn, accNumber string, limit int) (*[]models.Transaction, error) {
	connection, err := client.NewImmuClient(client.DefaultOptions())
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	response, err := connection.Login(
		ctx,
		[]byte(config.GetEnv("IMMUDB_USERNAME")),
		[]byte(config.GetEnv("IMMUDB_PASSWORD")),
	)
	if err != nil {
		return nil, err
	}
	md := metadata.Pairs("authorization", response.Token)
	ctx = metadata.NewOutgoingContext(ctx, md)
	result, err := connection.SQLQuery(
		ctx,
		fmt.Sprintf(`
		SELECT * FROM transactions
		WHERE account_number = @accNumber AND type = @type
		LIMIT %d`, limit),
		map[string]interface{}{"accNumber": accNumber, "type": _type.Value()},
		true,
	)
	if err != nil {
		return nil, err
	}
	var transactions []models.Transaction
	for _, row := range result.Rows {
		amountFloat, _ := strconv.ParseFloat(row.Values[2].GetS(), 64)
		txn := models.Transaction{
			TxnID:     row.Values[0].GetS(),
			Type:      row.Values[1].GetS(),
			Amount:    amountFloat,
			Number:    row.Values[3].GetS(),
			Timestamp: time.UnixMicro(row.Values[4].GetTs()),
		}
		transactions = append(transactions, txn)
	}
	return &transactions, nil
}
