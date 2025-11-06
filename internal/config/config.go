package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetToken reads TELE_TOKEN from .env or an environment variable
func GetToken() string {
	// Load .env (ignore error if file is missing)
	_ = godotenv.Load()

	token := os.Getenv("TELE_TOKEN")
	if token == "" {
		log.Fatal("❌ TELE_TOKEN is not set. Please check your .env file or environment variable.")
	}
	return token
}
