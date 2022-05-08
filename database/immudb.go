package database

import (
	"context"
	"time"

	"github.com/Devansh3712/go-banking-api/config"
	"github.com/Devansh3712/go-banking-api/utils"
	"github.com/codenotary/immudb/pkg/client"
	"google.golang.org/grpc/metadata"
)

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
		[]byte(config.EnvValue("IMMUDB_USERNAME")),
		[]byte(config.EnvValue("IMMUDB_PASSWORD")),
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

func CreateTransaction(_type string, amount string, accNumber string) (*string, error) {
	connection, err := client.NewImmuClient(client.DefaultOptions())
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	response, err := connection.Login(
		ctx,
		[]byte(config.EnvValue("IMMUDB_USERNAME")),
		[]byte(config.EnvValue("IMMUDB_PASSWORD")),
	)
	if err != nil {
		return nil, err
	}
	md := metadata.Pairs("authorization", response.Token)
	ctx = metadata.NewOutgoingContext(ctx, md)
	txnID, err := utils.GenerateTxnID()
	if err != nil {
		return nil, err
	}
	_, err = connection.SQLExec(
		ctx,
		`INSERT INTO transactions (id, type, amount, account_number, time)
		VALUES (@id, @type, @amount, @accNumber, @currTime)`,
		map[string]interface{}{
			"id":        *txnID,
			"type":      _type,
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
