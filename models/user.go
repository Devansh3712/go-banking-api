package models

import (
	"encoding/json"
	"time"
)

type User struct {
	Email     string    `json:"email,omitempty" gorm:"PRIMARY_KEY"`
	Password  string    `json:"password,omitempty" gorm:"NOT_NULL"`
	Phone     string    `json:"phone,omitempty" gorm:"NOT_NULL"`
	DOB       string    `json:"dob,omitempty" gorm:"NOT_NULL"`
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
		"dob":        user.DOB,
		"created_at": user.Timestamp,
	}
	return json.Marshal(data)
}
