package database

import (
	"context"
	"fmt"

	"github.com/Devansh3712/go-banking-api/config"
	"github.com/Devansh3712/go-banking-api/models"
	immudb "github.com/codenotary/immudb/pkg/client"
	"github.com/codenotary/immudb/pkg/stdlib"
)

func createOptions() *immudb.Options {
	options := immudb.DefaultOptions()
	options.Username = config.EnvValue("IMMUDB_USERNAME")
	options.Password = config.EnvValue("IMMUDB_PASSWORD")
	options.Database = config.EnvValue("IMMUDB_DATABASE")
	return options
}

func CreateAccount(account *models.Account) error {
	db := stdlib.OpenDB(createOptions())
	defer db.Close()
	_, err := db.ExecContext(
		context.TODO(),
		fmt.Sprintf("CREATE TABLE %s(type VARCHAR(10), amount INTEGER, timestamp VARCHAR(30))", account.Number),
	)
	if err != nil {
		return err
	}
	return nil
}
