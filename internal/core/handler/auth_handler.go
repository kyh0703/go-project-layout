package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kyh0703/layout/internal/core/middleware"
)

type AuthHandler interface {
	Handler
	SignUp(c *fiber.Ctx) error
	SignIn(c *fiber.Ctx) error
	SignOut(c *fiber.Ctx) error
	Refresh(c *fiber.Ctx) error
}

type authHandler struct {
	authMiddleware middleware.AuthMiddleware
}

func NewAuthHandler(
	authMiddleware middleware.AuthMiddleware,
) AuthHandler {
	return &authHandler{
		authMiddleware: authMiddleware,
	}
}

func (h *authHandler) Table() []Mapper {
	return []Mapper{
		Mapping(fiber.MethodGet, "/auth/whoami", h.authMiddleware.CurrentUser(), h.Whoami),
		Mapping(fiber.MethodPost, "/auth/signup", h.SignUp),
		Mapping(fiber.MethodPost, "/auth/signin", h.SignIn),
		Mapping(fiber.MethodPost, "/auth/signout", h.SignOut),
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
