package models

import "encoding/json"

type Account struct {
	Email  string  `gorm:"PRIMARY_KEY"`
	User   *User   `gorm:"FOREIGN_KEY:Email"`
	Amount float32 `gorm:"default:0"`
	Number string
}

// Remove email field from Account struct output.
func (account *Account) MarshalJSON() ([]byte, error) {
	data := map[string]any{
		"user":           account.User,
		"amount":         account.Amount,
		"account_number": account.Number,
	}
	return json.Marshal(data)
}
