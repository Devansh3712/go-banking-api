package database

import (
	"context"

	"github.com/Devansh3712/go-banking-api/config"
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
			type 			VARCHAR[10],
			amount 			VARCHAR,
			account_number 	INTEGER,
			PRIMARY KEY (id)
		)`,
		map[string]interface{}{},
	)
	if err != nil {
		return err
	}
	return nil
}
