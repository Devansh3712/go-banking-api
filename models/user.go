package models

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type User struct {
	Email     string    `json:"email" gorm:"primaryKey"`
	Password  string    `json:"-"`
	Phone     string    `json:"phone"`
	Timestamp time.Time `json:"created_at"`
}

type UserAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) Hash() error {
	hash := sha256.New()
	_, err := hash.Write([]byte(user.Password))
	if err != nil {
		return err
	}
	user.Password = fmt.Sprintf("%x", hash.Sum(nil))
	return nil
}

func (user *UserAuth) Hash() (*string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(user.Password))
	if err != nil {
		return nil, err
	}
	result := fmt.Sprintf("%x", hash.Sum(nil))
	return &result, nil
}
