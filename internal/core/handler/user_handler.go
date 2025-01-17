package handler

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	Handler
	CreateOne(c *fiber.Ctx) error
	GetOne(c *fiber.Ctx) error
	UpdateOne(c *fiber.Ctx) error
	DeleteOne(c *fiber.Ctx) error
}

type userHandler struct{}

func NewUserHandler() UserHandler {
	return &userHandler{}
}

func (u *userHandler) Table() []Mapping {
	return []Mapping{
		{Method: fiber.MethodPost, Path: "/user", Handler: u.CreateOne},
		{Method: fiber.MethodGet, Path: "/user/:id", Handler: u.GetOne},
		{Method: fiber.MethodPut, Path: "/user/:id", Handler: u.UpdateOne},
		{Method: fiber.MethodDelete, Path: "/user/:id", Handler: u.DeleteOne},
	}
}

func (u *userHandler) CreateOne(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (u *userHandler) GetOne(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (u *userHandler) UpdateOne(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (u *userHandler) DeleteOne(c *fiber.Ctx) error {
	panic("unimplemented")
}
