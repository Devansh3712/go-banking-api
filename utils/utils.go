package utils

import (
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
