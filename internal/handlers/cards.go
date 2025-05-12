package handlers

import (
	"fmt"
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

// @Summary Создание новой карточки
// @Description Создает новую карточку с английским словом и его переводом в указанной категории
// @Tags cards
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param card body models.Card true "Данные карточки"
// @Success 201 {object} models.Card "Карточка успешно создана"
// @Failure 400 {object} map[string]string "Неверные данные"
// @Failure 401 {object} map[string]string "Требуется авторизация"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /cards [post]
func CreateCard(c echo.Context) error {
	userID := c.Get("userID").(int)
	req := new(models.Card)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "неверные данные"})
	}

	fmt.Printf("Creating card for user %d in category %d\n", userID, req.CategoryID)

	// Проверяем, существует ли категория и принадлежит ли она пользователю
	var categoryExists bool
	db := db.GetDB()
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM categories WHERE id = $1 AND user_id = $2)",
		req.CategoryID, userID).Scan(&categoryExists)
	if err != nil {
		fmt.Printf("Error checking category: %v\n", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "ошибка при проверке категории"})
	}

	fmt.Printf("Category exists: %v\n", categoryExists)

	if !categoryExists {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "категория не найдена"})
	}

	// Создаем карточку
	card := models.Card{
		UserID:      userID,
		CategoryID:  req.CategoryID,
		Word:        req.Word,
		Translation: req.Translation,
	}

	err = db.QueryRow(
		"INSERT INTO cards(user_id, category_id, word, translation) VALUES($1, $2, $3, $4) RETURNING id",
		card.UserID, card.CategoryID, card.Word, card.Translation,
	).Scan(&card.ID)

	if err != nil {
		fmt.Printf("Error creating card: %v\n", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "не удалось создать карточку"})
	}

	return c.JSON(http.StatusCreated, card)
}

// @Summary Получение карточки по ID
// @Description Возвращает карточку по её ID
// @Tags cards
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "ID карточки"
// @Success 200 {object} models.Card "Карточка"
// @Failure 401 {object} map[string]string "Требуется авторизация"
// @Failure 404 {object} map[string]string "Карточка не найдена"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /cards/{id} [get]
func GetCard(c echo.Context) error {
	userID := c.Get("userID").(int)
	cardID := c.Param("id")

	var card models.Card
	db := db.GetDB()

	err := db.QueryRow(
		"SELECT id, user_id, category_id, word, translation FROM cards WHERE id = $1 AND user_id = $2",
		cardID, userID,
	).Scan(&card.ID, &card.UserID, &card.CategoryID, &card.Word, &card.Translation)

	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "карточка не найдена"})
	}

	return c.JSON(http.StatusOK, card)
}
