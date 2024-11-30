package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func NewConfigs() {
	if getEnv("APP_ENV", "production") == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Failed to load .env file")
		}
	}

	App = NewAppConfigs()
	Secret = NewSecretConfigs()
	DB = NewDBConfigs()
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		parseValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return parseValue
	}
	return fallback
}

func getEnvAsBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		parseValue, err := strconv.ParseBool(value)
		if err != nil {
			return fallback
		}
		return parseValue
	}
	return fallback
}
