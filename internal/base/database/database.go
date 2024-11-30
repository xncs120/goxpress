package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"goxpress/internal/base/config"
)

type Pgsql struct {
	db *gorm.DB
}

func NewPgsql() *Pgsql {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DB.Host, config.DB.Username, config.DB.Password, config.DB.Database, config.DB.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return &Pgsql{db: db}
}

func (d *Pgsql) GetDB() *gorm.DB {
	return d.db
}
