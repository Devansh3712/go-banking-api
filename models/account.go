package models

import (
	"encoding/json"
	"time"
)

type Account struct {
	Email   string  `gorm:"PRIMARY_KEY"`
	User    *User   `gorm:"FOREIGNKEY:Email"`
	Balance float32 `gorm:"default:100"`
	Number  string
}

type Transaction struct {
	TxnID     string
	Type      string
	Amount    string
	Number    string
	Timestamp time.Time
}

// Remove email field from Account struct output.
func (account *Account) MarshalJSON() ([]byte, error) {
	data := map[string]any{
		"user":           account.User,
		"balance":        account.Balance,
		"account_number": account.Number,
	}
	return json.Marshal(data)
}

func (transaction *Transaction) MarshalJSON() ([]byte, error) {
	data := map[string]any{
		"transaction_id":   transaction.TxnID,
		"transaction_type": transaction.Type,
		"amount":           transaction.Amount,
		"timestamp":        transaction.Timestamp,
	}
	return json.Marshal(data)
}
