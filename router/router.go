package router

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type Router struct {
	e  *echo.Echo
	db *gorm.DB
}

func New(e *echo.Echo, db *gorm.DB) *Router {
	return &Router{
		e:  e,
		db: db,
	}
}
