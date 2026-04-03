package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func New() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		parseValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return parseValue
	}
	return fallback
}

func GetEnvAsBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		parseValue, err := strconv.ParseBool(value)
		if err != nil {
			return fallback
		}
		return parseValue
	}
	return fallback
}
