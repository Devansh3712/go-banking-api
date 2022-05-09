package models

// Enum for adding transaction type in ImmuDB.
type Txn string

const (
	Deposit  Txn = "deposit"
	Withdraw Txn = "withdraw"
)

func (t Txn) Value() string {
	return string(t)
}
