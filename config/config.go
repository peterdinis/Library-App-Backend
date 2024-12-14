package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config function to get environment value by key
func Config(key string) string {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}

	// Return the environment variable value for the given key
	return os.Getenv(key)
}