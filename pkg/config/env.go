package config

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvVar string // EnvVar represents an environment variable name.

// Server configuration.
const (
	ServerPortVar EnvVar = "SERVER_PORT"
)

// Database configuration.
const (
	DatabaseHostVar     EnvVar = "DATABASE_HOST"
	DatabaseUserVar     EnvVar = "DATABASE_USER"
	DatabasePasswordVar EnvVar = "DATABASE_PASSWORD"
	DatabaseNameVar     EnvVar = "DATABASE_NAME"
	DatabasePortVar     EnvVar = "DATABASE_PORT"
	DatabaseSSlModeVar  EnvVar = "DATABASE_SSLMODE"
	DatabaseTimeZoneVar EnvVar = "DATABASE_TIMEZONE"
)

// Database represents the database configuration.
type Database struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSlMode  string
	TimeZone string
}

// NewDatabase creates a new DatabaseConfig.
func NewDatabase() *Database {
	return &Database{
		Host:     GetEnv(DatabaseHostVar),
		User:     GetEnv(DatabaseUserVar),
		Password: GetEnv(DatabasePasswordVar),
		DBName:   GetEnv(DatabaseNameVar),
		Port:     GetEnv(DatabasePortVar),
		SSlMode:  GetEnv(DatabaseSSlModeVar),
		TimeZone: GetEnv(DatabaseTimeZoneVar),
	}
}

// getEnv returns the value of an environment variable.
func GetEnv(envVar EnvVar) string {
	return os.Getenv(string(envVar))
}

// loadEnv loads the environment variables from the .env file.
func LoadEnv() error {
	return godotenv.Load()
}
