package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Hash passwords using SHA256.
func Hash(password string) (*string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(password))
	if err != nil {
		return nil, err
	}
	result := fmt.Sprintf("%x", hash.Sum(nil))
	return &result, nil
}

// Create a random transaction ID.
func GenerateTxnID() (*string, error) {
	bytes := make([]byte, 8)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	id := fmt.Sprintf("%x", bytes)
	return &id, nil
}

// Get environment variable values.
func GetEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	return os.Getenv(key)
}
