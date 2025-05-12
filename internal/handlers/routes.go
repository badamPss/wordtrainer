package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"wordtrainer/internal/config"
	"wordtrainer/internal/middleware"
)

func RegisterRoutes(e *echo.Echo, cfg *config.Config) {
	fmt.Println("ZZZZZZZZZZZZ")
	e.POST("/register", Register)
	e.POST("/login", Login)

	// Защищённые маршруты
	e.GET("/categories", GetCategories, middleware.JWTMiddleware(cfg))
	e.GET("/cards", GetCards, middleware.JWTMiddleware(cfg))
}
