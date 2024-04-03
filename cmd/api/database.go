package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
	port     = "5432"
	sslmode  = "disable"
	timeZone = "Asia/Shanghai"
)

type TB01 struct {
	ID       uint
	ColTexto string
	ColDt    time.Time
}

func dsn() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timeZone,
	)
}

func (t *TB01) Create() error {
	db, err := gorm.Open(postgres.Open(dsn()), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	t.ColDt = time.Now()
	if err := db.Create(t).Error; err != nil {
		return fmt.Errorf("failed to insert data on database: %w", err)
	}

	return nil
}
