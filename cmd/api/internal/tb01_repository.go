package internal

import (
	"circle/pkg/database"
	"circle/pkg/tb01"
	"fmt"
)

// TB01Repository represents a GORM repository for TB01.
type TB01Repository struct {
	db *database.PostgreSQL
}

// NewTB01Repository creates a new TB01GormRepository.
func NewTB01Repository(db *database.PostgreSQL) *TB01Repository {
	return &TB01Repository{db}
}

// Create inserts a new record into the database.
func (r *TB01Repository) Create(t *tb01.TB01) error {
	if err := r.db.Create(t).Error; err != nil {
		return fmt.Errorf("failed to insert data on database: %w", err)
	}
	return nil
}
