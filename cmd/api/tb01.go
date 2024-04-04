package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TB01 represents a record in the database.
type TB01 struct {
	ID       uint      `gorm:"primaryKey" json:"id,omitempty"`
	ColTexto string    `json:"col_texto,omitempty"`
	ColDt    time.Time `json:"col_dt,omitempty"`
}

// TB01Service represents a service for TB01.
type TB01Service interface {
	Create(*TB01) error
}

// TB01DefaultService represents a service for TB01.
type TB01DefaultService struct {
	repo TB01Repository
}

// NewTB01DefaultService creates a new TB01DefaultService.
func NewTB01DefaultService(repo TB01Repository) *TB01DefaultService {
	return &TB01DefaultService{repo}
}

// Create inserts a new record into the database.
func (s *TB01DefaultService) Create(t *TB01) error {
	t.ColDt = time.Now()
	return s.repo.Create(t)
}

// TB01Repository represents a repository for TB01.
type TB01Repository interface {
	Create(*TB01) error
}

// TB01GormRepository represents a GORM repository for TB01.
type TB01GormRepository struct {
	db *Database
}

// NewTB01GormRepository creates a new TB01GormRepository.
func NewTB01GormRepository(db *Database) *TB01GormRepository {
	return &TB01GormRepository{db}
}

// Create inserts a new record into the database.
func (r *TB01GormRepository) Create(t *TB01) error {
	if err := r.db.Create(t).Error; err != nil {
		return fmt.Errorf("failed to insert data on database: %w", err)
	}
	return nil
}

// Database represents a database connection.
type Database struct {
	*gorm.DB
}

// NewDatabase creates a new Database.
func NewPostgresRepository(cfg *DatabaseConfig) (*Database, error) {
	db, err := gorm.Open(postgres.Open(dsn()), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}
	return &Database{db}, nil
}
