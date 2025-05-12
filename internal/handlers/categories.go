package handlers

import (
	"net/http"
	"wordtrainer/internal/db"
	"wordtrainer/internal/models"

	"github.com/labstack/echo/v4"
)

// @Summary Создание новой категории
// @Description Создает новую категорию для пользователя
// @Tags categories
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param category body models.Category true "Данные категории"
// @Success 201 {object} models.Category "Категория успешно создана"
// @Failure 400 {object} map[string]string "Неверные данные"
// @Failure 401 {object} map[string]string "Требуется авторизация"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /categories [post]
func CreateCategory(c echo.Context) error {
	userID := c.Get("userID").(int)
	req := new(models.Category)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "неверные данные"})
	}

	category := models.Category{
		UserID: userID,
		Name:   req.Name,
	}

	db := db.GetDB()
	err := db.QueryRow(
		"INSERT INTO categories(user_id, name) VALUES($1, $2) RETURNING id",
		category.UserID, category.Name,
	).Scan(&category.ID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "не удалось создать категорию"})
	}

	return c.JSON(http.StatusCreated, category)
}

// @Summary Получение списка категорий
// @Description Возвращает список всех категорий пользователя
// @Tags categories
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} models.Category "Список категорий"
// @Failure 401 {object} map[string]string "Требуется авторизация"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /categories [get]
func GetCategories(c echo.Context) error {
	// Получаем userID из контекста (он был установлен в JWT Middleware)
	userID := c.Get("userID").(int)

	// Используем глобальную базу данных
	db := db.GetDB()

	var categories []models.Category

	// Получаем все категории пользователя
	err := db.Select(&categories, "SELECT id, name FROM categories WHERE user_id = $1", userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "не удалось получить категории"})
	}

	return c.JSON(http.StatusOK, categories)
}
