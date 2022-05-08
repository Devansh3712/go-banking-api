package models

import "encoding/json"

type Account struct {
	Email   string  `gorm:"PRIMARY_KEY"`
	User    *User   `gorm:"FOREIGNKEY:Email"`
	Balance float32 `gorm:"default:100"`
	Number  string
}

// Remove email field from Account struct output.
func (account *Account) MarshalJSON() ([]byte, error) {
	data := map[string]any{
		"user":           account.User,
		"amount":         account.Balance,
		"account_number": account.Number,
	}
	return json.Marshal(data)
}
