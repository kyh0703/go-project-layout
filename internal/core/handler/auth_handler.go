package handler

import (
	"github.com/gofiber/fiber/v2"
)

type authHandler struct{}

func NewAuthHandler() Handler {
	return &authHandler{}
}

func (h *authHandler) Table() []Mapping {
	return []Mapping{
		{Method: fiber.MethodGet, Path: "/auth/whoami", Handler: h.Whoami},
		{Method: fiber.MethodPost, Path: "/auth/signup", Handler: h.SignUp},
		{Method: fiber.MethodPost, Path: "/auth/signin", Handler: h.SignIn},
		{Method: fiber.MethodPost, Path: "/auth/signout", Handler: h.SignOut},
	}
}

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
