package auth

import (
	"github.com/gofiber/fiber/v2"
)

type authHandler struct{}

func (h *authHandler) Whoami(c *fiber.Ctx) error {
	user := c.Locals("user")
	return c.JSON(user)
}

func (h *authHandler) SignUp(c *fiber.Ctx) error {
	return nil
}

func (h *authHandler) SignIn(c *fiber.Ctx) error {
	return nil
}

func (h *authHandler) SignOut(c *fiber.Ctx) error {
	return nil
}

func (h *authHandler) Refresh(c *fiber.Ctx) error {
	return nil
}
