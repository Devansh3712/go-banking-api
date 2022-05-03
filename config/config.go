package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Get environment variable values.
func EnvValue(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	return os.Getenv(key)
}
