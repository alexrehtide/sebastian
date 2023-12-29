package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func New() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
		return nil, err
	}

	return &Config{
		PostgresUser:     getEnv("POSTGRES_USER", "postgres"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", "3769"),
		PostgresName:     getEnv("POSTGRES_NAME", "postgres"),
		PostgresHost:     getEnv("POSTGRES_HOST", "postgres"),
		PostgresPort:     getEnv("POSTGRES_PORT", "5432"),
		Debug:            getEnv("DEBUG", "false"),
	}, nil
}

type Config struct {
	PostgresHost     string
	PostgresUser     string
	PostgresPassword string
	PostgresName     string
	PostgresPort     string
	Debug            string
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
