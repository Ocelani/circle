package tb01

import (
	"time"
)

// TB01 represents a record in the database.
type TB01 struct {
	ID       uint      `gorm:"primaryKey" json:"id,omitempty"`
	ColTexto string    `json:"col_texto,omitempty"`
	ColDt    time.Time `json:"col_dt,omitempty"`
}

// Service represents a service for TB01.
type Service interface {
	Create(*TB01) error
}

// Repository represents a repository for TB01.
type Repository interface {
	Create(*TB01) error
}
