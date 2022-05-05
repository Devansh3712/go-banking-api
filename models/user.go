package models

import (
	"encoding/json"
	"time"
)

type User struct {
	Email     string    `json:"email" gorm:"PRIMARY_KEY"`
	Password  string    `json:"password" gorm:"NOT_NULL"`
	Phone     string    `json:"phone" gorm:"NOT_NULL"`
	Timestamp time.Time `json:"created_at"`
}

type UserAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Remove password field from User struct output.
func (user *User) MarshalJSON() ([]byte, error) {
	data := map[string]any{
		"email":      user.Email,
		"phone":      user.Phone,
		"created_at": user.Timestamp,
	}
	return json.Marshal(data)
}
