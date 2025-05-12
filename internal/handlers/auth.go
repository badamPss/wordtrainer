package handlers

import (
	"net/http"
	"wordtrainer/internal/db"
	"wordtrainer/internal/models"
	"wordtrainer/internal/utils"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// @Summary Регистрация нового пользователя
// @Description Создает нового пользователя в системе
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "Данные пользователя"
// @Success 201 {object} map[string]interface{} "Пользователь успешно создан"
// @Failure 400 {object} map[string]string "Неверные данные"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /register [post]
func Register(c echo.Context) error {
	req := new(models.User)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "неверные данные"})
	}

	// Хэширование пароля
	hash, err := bcrypt.GenerateFromPassword([]byte(req.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "ошибка при хэшировании пароля"})
	}

	// Создание пользователя
	user := models.User{
		Username:     req.Username,
		PasswordHash: string(hash),
	}

	// Получаем подключение к базе данных
	db := db.GetDB()

	// Вставка пользователя в базу данных
	if err := db.QueryRow("INSERT INTO users(username, password_hash) VALUES($1, $2) RETURNING id", user.Username, user.PasswordHash).Scan(&user.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "не удалось создать пользователя"})
	}

	return c.JSON(http.StatusCreated, echo.Map{"id": user.ID, "username": user.Username})
}

// @Summary Вход в систему
// @Description Аутентификация пользователя и получение JWT токена
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "Данные для входа"
// @Success 200 {object} map[string]string "JWT токен"
// @Failure 400 {object} map[string]string "Неверные данные"
// @Failure 401 {object} map[string]string "Неверные учетные данные"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /login [post]
func Login(c echo.Context) error {
	req := new(models.User)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "неверные данные"})
	}

	// Поиск пользователя по имени
	var user models.User
	db := db.GetDB()

	err := db.QueryRow("SELECT id, username, password_hash FROM users WHERE username = $1", req.Username).Scan(&user.ID, &user.Username, &user.PasswordHash)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "неверные данные"})
	}

	// Проверка пароля
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.PasswordHash))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "неверные данные"})
	}

	// Генерация JWT
	token, err := utils.GenerateJWT(user.ID, "supersecretkey")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "ошибка генерации токена"})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}
