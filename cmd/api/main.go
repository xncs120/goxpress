package main

import (
	"log"

	"github.com/labstack/echo/v4"

	"goxpress/internal/base/config"
	"goxpress/internal/base/database"
	"goxpress/internal/base/resource"
	"goxpress/internal/base/routes"
)

func main() {
	config.NewConfigs()

	e := echo.New()

	newPgsql := database.NewPgsql()
	db := newPgsql.GetDB()

	e.GET("/statics/*", resource.StaticHandler())
	templateRenderer := resource.NewTemplateRenderer()
	e.Renderer = templateRenderer

	newRoutes := routes.NewRoutes(e, db)
	newRoutes.RegisterPaths()

	port := config.App.Port
	if err := e.Start(":" + port); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
