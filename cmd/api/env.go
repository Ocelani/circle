package main

import (
	"circle/pkg/database"
	"os"
)

// EnvVar represents an environment variable name.
type EnvVar string

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

// NewDatabaseConfig creates a new DatabaseConfig.
func NewDatabaseConfig() *database.Config {
	return &database.Config{
		Host:     GetEnv(DatabaseHostVar),
		User:     GetEnv(DatabaseUserVar),
		Password: GetEnv(DatabasePasswordVar),
		DBName:   GetEnv(DatabaseNameVar),
		Port:     GetEnv(DatabasePortVar),
		SSlMode:  GetEnv(DatabaseSSlModeVar),
		TimeZone: GetEnv(DatabaseTimeZoneVar),
	}
}

// GetEnv returns the value of an environment variable.
func GetEnv(envVar EnvVar) string {
	return os.Getenv(string(envVar))
}

// GetPort returns the server port.
func GetPort() string {
	port := GetEnv(ServerPortVar)
	if port == "" {
		return ":" + DefaultPort
	}
	return ":" + port
}
