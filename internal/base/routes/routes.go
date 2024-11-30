package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/xncs120/goxpress/internal/base/auth"
	"github.com/xncs120/goxpress/internal/landing"
	"github.com/xncs120/goxpress/internal/user"
)

type Routes struct {
	e  *echo.Echo
	db *gorm.DB
}

func NewRoutes(e *echo.Echo, db *gorm.DB) *Routes {
	return &Routes{
		e:  e,
		db: db,
	}
}

func (r *Routes) RegisterPaths() {
	landingHandler := landing.NewHandler()
	userHandler := user.NewHandler(r.db)

	// web
	r.e.GET("/", landingHandler.Index)
	r.e.GET("/register", landingHandler.Register)
	r.e.GET("/docs", landingHandler.Docs)

	// api
	r.e.POST("/api/register", userHandler.CreateUser)
	r.e.POST("/api/token", userHandler.GenerateToken)

	api := r.e.Group("/api/v1")
	api.Use(auth.JWTAuthorization())

	userApi := api.Group("/users")
	userApi.GET("/:id", userHandler.GetUser)
	userApi.PUT("/:id", userHandler.UpdateUser)
}
