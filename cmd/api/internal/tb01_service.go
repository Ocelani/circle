package internal

import (
	"circle/pkg/tb01"
	"fmt"
	"time"
)

// TB01Service represents a service for TB01.
type TB01Service struct {
	repo tb01.Repository
}

// NewTB01Service creates a new TB01DefaultService.
func NewTB01Service(repo tb01.Repository) *TB01Service {
	return &TB01Service{
		repo: repo,
	}
}

// Create inserts a new record into the database.
func (s *TB01Service) Create(data *tb01.TB01) error {
	if data.ColTexto == "" {
		return fmt.Errorf("empty ColTexto field")
	}
	data.ColDt = time.Now()

	return s.repo.Create(data)
}
