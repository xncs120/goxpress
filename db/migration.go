package db

import (
	"github.com/xncs120/goxpress/models"
)

func (db *DB) Migration() {
	db.Gorm.AutoMigrate(
		&models.User{},
	)
}
