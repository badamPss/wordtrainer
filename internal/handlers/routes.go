package handlers

import (
	"wordtrainer/internal/config"
	"wordtrainer/internal/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, cfg *config.Config) {
	e.POST("/register", Register)
	e.POST("/login", Login)

	// Защищённые маршруты
	e.GET("/categories", GetCategories, middleware.JWTMiddleware(cfg))
	e.POST("/categories", CreateCategory, middleware.JWTMiddleware(cfg))
	e.GET("/cards", GetCards, middleware.JWTMiddleware(cfg))
	e.POST("/cards", CreateCard, middleware.JWTMiddleware(cfg))
	e.GET("/cards/:id", GetCard, middleware.JWTMiddleware(cfg))
}
