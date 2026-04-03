package main

import (
	"log"

	"goxpress/config"
	"goxpress/db"
	"goxpress/internal/request"
	"goxpress/router"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

func main() {
	config.New()

	db := db.New()
	// db.Migration() // remove if manual migration is preferred

	e := echo.New()

	e.Validator = &request.CustomValidator{Validator: validator.New()}

	router := router.New(e, db.Gorm)
	router.AddWebRoutes()
	router.AddApiRoutes()

	if err := e.Start(":" + config.App.Port); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
