package middleware

import (
	"strings"

	"powerapp/server/database"
	"powerapp/server/models"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{"error": "Missing authorization header"})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		var authToken models.AuthToken
		if err := database.DB.Preload("User").Where("token = ?", tokenString).First(&authToken).Error; err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
		}

		if authToken.IsExpired() {
			database.DB.Delete(&authToken)
			return c.Status(401).JSON(fiber.Map{"error": "Token expired"})
		}

		c.Locals("user_id", authToken.UserID)
		c.Locals("user", authToken.User)
		c.Locals("auth_token", &authToken)
		return c.Next()
	}
}
