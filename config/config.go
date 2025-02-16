package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Println("no .env file")
	}
}

func GetEnv(key, defaultValue string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}

	return defaultValue
}
