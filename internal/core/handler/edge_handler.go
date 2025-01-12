package handler

import "github.com/gofiber/fiber/v2"

type EdgeHandler interface {
	Handler
	CreateOne(c *fiber.Ctx) error
	UpdateOne(c *fiber.Ctx) error
	DeleteOne(c *fiber.Ctx) error
}

type edgeHandler struct{}

func NewEdgeHandler() EdgeHandler {
	return &edgeHandler{}
}

func (h *edgeHandler) Setup(router fiber.Router) {
	router.Post("/edge", h.CreateOne)
	router.Put("/edge/:id", h.UpdateOne)
	router.Delete("/edge/:id", h.DeleteOne)
}

func (h *edgeHandler) CreateOne(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *edgeHandler) DeleteOne(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *edgeHandler) UpdateOne(c *fiber.Ctx) error {
	panic("unimplemented")
}
