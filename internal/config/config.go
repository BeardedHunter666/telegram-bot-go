package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetToken зчитує TELE_TOKEN із .env або змінної середовища
func GetToken() string {
	// Завантажуємо .env (ігноруємо помилку, якщо файл відсутній)
	_ = godotenv.Load()

	token := os.Getenv("TELE_TOKEN")
	if token == "" {
		log.Fatal("❌ TELE_TOKEN is not set. Please check your .env file or environment variable.")
	}
	return token
}
