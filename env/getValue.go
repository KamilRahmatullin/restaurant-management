package env

import (
	"os"

	"github.com/joho/godotenv"
)

func GetValue(key, fallback string) string {

	if err := godotenv.Load(); err != nil {
		return fallback
	}

	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}
