package auth

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kyh0703/go-project-layout/internal/auth/dto"
)

type AuthController struct {
	validator *validator.Validate
}

func (ctrl *AuthController) SetupRoutes(router fiber.Router) fiber.Router {
	auth := router.Group("/user")
	auth.Get("/whoami", ctrl.Whoami)
	auth.Post("/signup", ctrl.Signup)
	auth.Post("/signin", ctrl.Signin)
	auth.Post("/signout", ctrl.Signout)
	auth.Post("/refresh", ctrl.Signin)
	return auth
}

func (ctrl *AuthController) Whoami(c *fiber.Ctx) error {
	user := c.Locals("user")
	return c.JSON(user)
}

func (ctrl *AuthController) Signup(c *fiber.Ctx) error {
	var signup dto.Signup
	if err := c.BodyParser(&signup); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	if err := ctrl.validator.StructCtx(c.Context(), signup); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(signup)
}

func (ctrl *AuthController) Signin(c *fiber.Ctx) error {
	var signin dto.Signin
	if err := c.BodyParser(&signin); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	if err := ctrl.validator.StructCtx(c.Context(), signin); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(signin)
}

func (ctrl *AuthController) Signout(c *fiber.Ctx) error {
	c.Locals("user", nil)
	return c.SendStatus(fiber.StatusNoContent)
}

func (ctrl *AuthController) Refresh(c *fiber.Ctx) error {
	return nil
}
