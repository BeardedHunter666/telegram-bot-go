package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetToken reads BOT_TOKEN from .env or an environment variable
func GetToken() string {
	// Load .env (ignore error if file is missing)
	_ = godotenv.Load()

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("❌ BOT_TOKEN is not set. Please check your .env file or environment variable.")
	}
	return token
}
