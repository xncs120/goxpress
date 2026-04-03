package db

import (
	"goxpress/models"
)

func (db *DB) Migration() {
	db.Gorm.AutoMigrate(
		&models.User{},
	)
}
