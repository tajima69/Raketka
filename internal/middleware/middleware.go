package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
	"time"
)

var SecretKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(userID int, username string) (string, error) {
	claims := jwt.MapClaims{
		"id":       userID,
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "missing or malformed token"})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "invalid or expired token"})
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Locals("userID", claims["id"])
		c.Locals("username", claims["username"])

		return c.Next()
	}
}
