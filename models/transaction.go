package models

import (
	"encoding/json"
	"time"
)

// Enum for adding transaction type in ImmuDB.
type Txn string

const (
	Deposit  Txn = "deposit"
	Withdraw Txn = "withdraw"
)

func (t Txn) Value() string {
	return string(t)
}

type Transaction struct {
	TxnID     string
	Type      string
	Amount    uint64
	Number    string
	Timestamp time.Time
}

// Remove account number from Transaction struct output.
func (transaction *Transaction) MarshalJSON() ([]byte, error) {
	data := map[string]any{
		"transaction_id":   transaction.TxnID,
		"transaction_type": transaction.Type,
		"amount":           transaction.Amount,
		"timestamp":        transaction.Timestamp,
	}
	return json.Marshal(data)
}
