package router

import (
	"goxpress/handlers"
	"goxpress/internal/security"
)

func (r *Router) AddApiRoutes() {
	authHandler := handlers.NewAuthHandler(r.db)
	userHandler := handlers.NewUserHandler(r.db)

	api := r.e.Group("/api")
	api.GET("/health", authHandler.Health)
	api.POST("/register", authHandler.Register)
	api.POST("/token", authHandler.Login)

	apiV1 := r.e.Group("/api/v1")
	apiV1.Use(security.JWTAuthorization())

	userApi := apiV1.Group("/users")
	userApi.GET("/:id", userHandler.GetUser)
	userApi.PATCH("/:id", userHandler.UpdateUser)
}
