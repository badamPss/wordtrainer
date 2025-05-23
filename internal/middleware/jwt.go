package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"wordtrainer/internal/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(cfg *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			fmt.Printf("Auth header: %s\n", authHeader)

			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "требуется токен"})
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "неверный формат токена"})
			}

			tokenStr := parts[1]
			token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.ErrUnauthorized
				}
				return []byte(cfg.JWTSecret), nil
			})
			if err != nil || !token.Valid {
				fmt.Printf("Token validation error: %v\n", err)
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "недействительный токен"})
			}

			claims := token.Claims.(jwt.MapClaims)
			userID := int(claims["user_id"].(float64))
			fmt.Printf("User ID from token: %d\n", userID)
			c.Set("userID", userID)

			return next(c)
		}
	}
}
