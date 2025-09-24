package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

func LoadEnv(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open .env file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Split on first '='
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Set environment variables
		os.Setenv(key, value)
	}

	return scanner.Err()
}

// Retrieves database configuration from env variables
func GetDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
	}
}
