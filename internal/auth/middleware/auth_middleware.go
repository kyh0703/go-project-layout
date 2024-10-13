package middleware

import "github.com/gofiber/fiber/v2"

type AuthMiddleware struct{}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func CurrentUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}
