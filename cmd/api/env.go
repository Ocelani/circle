package main

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	EnvVar string // EnvVar represents an environment variable name.
)

// Database configuration.
const (
	databaseHostVar     EnvVar = "DATABASE_HOST"
	databaseUserVar     EnvVar = "DATABASE_USER"
	databasePasswordVar EnvVar = "DATABASE_PASSWORD"
	databaseNameVar     EnvVar = "DATABASE_NAME"
	databasePortVar     EnvVar = "DATABASE_PORT"
	databaseSSlModeVar  EnvVar = "DATABASE_SSLMODE"
	databaseTimeZoneVar EnvVar = "DATABASE_TIMEZONE"
)

// getEnv returns the value of an environment variable.
func getEnv(envVar EnvVar) string {
	return os.Getenv(string(envVar))
}

// loadEnv loads the environment variables from the .env file.
func loadEnv() error {
	return godotenv.Load()
}
