package database

import (
	"circle/pkg/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config represents the database configuration.
type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSlMode  string
	TimeZone string
}

// PostgreSQL represents a database connection.
type PostgreSQL struct {
	*gorm.DB
}

// NewPostgreSQL creates a new Database.
func NewPostgreSQL(cfg *config.Database) (*PostgreSQL, error) {
	pg := postgres.Open(dsn(cfg))

	db, err := gorm.Open(pg, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	return &PostgreSQL{db}, nil
}

// dsn returns the data source name for the database connection.
func dsn(cfg *config.Database) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSlMode, cfg.TimeZone,
	)
}
