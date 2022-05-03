package models

import "encoding/json"

type Account struct {
	Email  string  `json:"email" gorm:"primaryKey"`
	User   *User   `json:"user" gorm:"foreignKey:Email"`
	Amount float32 `json:"amount" gorm:"default:0"`
}

// Remove email field from Account struct output.
func (account *Account) MarshalJSON() ([]byte, error) {
	data := map[string]any{
		"user":   account.User,
		"amount": account.Amount,
	}
	return json.Marshal(data)
}
