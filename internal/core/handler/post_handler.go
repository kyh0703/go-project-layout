package handler

import "github.com/gofiber/fiber/v2"

type PostHandler interface {
	Handler
	CreateOne(c *fiber.Ctx) error
	GetOne(c *fiber.Ctx) error
	DeleteOne(c *fiber.Ctx) error
	UpdateOne(c *fiber.Ctx) error
	Capture(c *fiber.Ctx) error
	Undo(c *fiber.Ctx) error
	Redo(c *fiber.Ctx) error
}

type postHandler struct{}

func NewPostHandler() PostHandler {
	return &postHandler{}
}

func (h *postHandler) Table() []Mapper {
	return []Mapper{}
}

func (h *postHandler) CreateOne(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *postHandler) GetOne(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *postHandler) DeleteOne(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *postHandler) UpdateOne(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *postHandler) Capture(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *postHandler) Undo(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *postHandler) Redo(c *fiber.Ctx) error {
	panic("unimplemented")
}
