package user

import "github.com/gofiber/fiber/v2"

type UserController struct{}

func (ctrl *UserController) SetupRoutes(router fiber.Router) fiber.Router {
	user := router.Group("/user")
	user.Get("/:id", ctrl.GetUser)
	user.Post("/", ctrl.CreateUser)
	user.Patch("/:id", ctrl.UpdateUser)
	user.Delete("/:id", ctrl.DeleteUser)
	return user
}

func (ctrl *UserController) GetUser(c *fiber.Ctx) error {
	return nil
}

func (ctrl *UserController) CreateUser(c *fiber.Ctx) error {
	return nil
}

func (ctrl *UserController) UpdateUser(c *fiber.Ctx) error {
	return nil
}

func (ctrl *UserController) DeleteUser(c *fiber.Ctx) error {
	return nil
}
