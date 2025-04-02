package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Port string
var ApiUrl string
var SigningKey string

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using defaults")
	}

	Port = getEnv("PORT", "8080")
	ApiUrl = getEnv("NEXT_PUBLIC_API_URL", "http://localhost:8080")
	SigningKey = getEnv("SIGNING_KEY", "")
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
