package handlers

import (
	"net/http"
	"wordtrainer/internal/db"
	"wordtrainer/internal/models"

	"github.com/labstack/echo/v4"
)

// @Summary Получение списка карточек
// @Description Возвращает список всех карточек пользователя
// @Tags cards
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} models.Card "Список карточек"
// @Failure 401 {object} map[string]string "Требуется авторизация"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /cards [get]
func GetCards(c echo.Context) error {
	// Получаем userID из контекста (он был установлен в JWT Middleware)
	userID := c.Get("userID").(int)

	// Используем глобальную базу данных
	db := db.GetDB()

	var cards []models.Card

	// Получаем все карточки пользователя
	err := db.Select(&cards, "SELECT id, word, translation, category_id FROM cards WHERE user_id = $1", userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "не удалось получить карточки"})
	}

	return c.JSON(http.StatusOK, cards)
}
