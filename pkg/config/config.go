package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
}

func LoadConfig() Config {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:root@localhost:5433/opengine?sslmode=disable"
	}

	return Config{
		DatabaseURL: dbURL,
	}
}
