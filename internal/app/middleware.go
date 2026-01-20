package app

import (
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/musishere/Blog/config"
	"github.com/musishere/Blog/internal/auth"
)

func InjectDatabase(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	}
}

func JWTMiddleware(cfg *config.AppConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Token is required"})
			}

			// Expect "Bearer <token>" format
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader { // No "Bearer " prefix
				return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid token format"})
			}

			claims, err := auth.ValidateJsonWebToken(tokenString)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, echo.Map{"mesage": "Unauthorized user"})
			}

			c.Set("user_id", claims.UserID)
			c.Set("role", claims.Role)
			return next(c)
		}
	}
}
