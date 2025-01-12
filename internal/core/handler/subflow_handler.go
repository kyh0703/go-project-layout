package handler

import "github.com/gofiber/fiber/v2"

type SubFlowHandler interface {
	Handler
	CreateOne(c *fiber.Ctx) error
	UpdateOne(c *fiber.Ctx) error
	DeleteOne(c *fiber.Ctx) error
	Capture(c *fiber.Ctx) error
	Undo(c *fiber.Ctx) error
	Redo(c *fiber.Ctx) error
}

type subFlowHandler struct{}

func NewSubFlowHandler() SubFlowHandler {
	return &subFlowHandler{}
}

func (h *subFlowHandler) Setup(router fiber.Router) {
	router.Post("/subflow", h.CreateOne)
	router.Put("/subflow/:id", h.UpdateOne)
	router.Delete("/subflow/:id", h.DeleteOne)
	router.Post("/subflow/:id/capture", h.Capture)
	router.Post("/subflow/:id/undo", h.Undo)
	router.Post("/subflow/:id/redo", h.Redo)
}

func (h *subFlowHandler) CreateOne(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *subFlowHandler) DeleteOne(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *subFlowHandler) UpdateOne(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *subFlowHandler) Capture(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *subFlowHandler) Undo(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *subFlowHandler) Redo(c *fiber.Ctx) error {
	panic("unimplemented")
}
