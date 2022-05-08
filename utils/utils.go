package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func Hash(password string) (*string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(password))
	if err != nil {
		return nil, err
	}
	result := fmt.Sprintf("%x", hash.Sum(nil))
	return &result, nil
}

func GenerateTxnID() (*string, error) {
	bytes := make([]byte, 8)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	id := fmt.Sprintf("%x", bytes)
	return &id, nil
}
