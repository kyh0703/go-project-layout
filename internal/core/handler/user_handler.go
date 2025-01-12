package handler

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	Handler
	GetOne(c *fiber.Ctx) error
	CreateOne(c *fiber.Ctx) error
	UpdateOne(c *fiber.Ctx) error
	DeleteOne(c *fiber.Ctx) error
}

type userHandler struct{}

func NewUserHandler() UserHandler {
	return &userHandler{}
}

func (u *userHandler) Setup(router fiber.Router) {
	panic("unimplemented")
}

func (u *userHandler) UpdateOne(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (u *userHandler) CreateOne(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (u *userHandler) DeleteOne(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (u *userHandler) GetOne(c *fiber.Ctx) error {
	panic("unimplemented")
}
