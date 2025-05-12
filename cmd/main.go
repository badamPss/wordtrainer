package main

import (
	"fmt"
	"log"
	_ "wordtrainer/docs" // Это будет создано после генерации Swagger
	"wordtrainer/internal/config"
	"wordtrainer/internal/db"
	"wordtrainer/internal/handlers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Word Trainer API
// @version 1.0
// @description API для приложения Word Trainer
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка при загрузке .env файла: ", err)
	}

	// Чтение конфигурации
	cfg := config.Load()
	fmt.Println(cfg)
	// Подключение к базе данных
	_, err = db.Connect(cfg)
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных: ", err)
	}

	// Создание Echo сервера
	e := echo.New()

	// Добавляем Swagger UI
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Регистрация маршрутов
	handlers.RegisterRoutes(e, cfg)

	// Запуск сервера
	e.Logger.Fatal(e.Start(":8080"))
}
