package handler

import "github.com/gofiber/fiber/v2"

type NodeHandler interface {
	Handler
	CreateOne(c *fiber.Ctx) error
	UpdateOne(c *fiber.Ctx) error
	DeleteOne(c *fiber.Ctx) error
}

type nodeHandler struct{}

func NewNodeHandler() NodeHandler {
	return &nodeHandler{}
}

func (h *nodeHandler) Setup(router fiber.Router) {
	router.Post("/node", h.CreateOne)
	router.Put("/node/:id", h.UpdateOne)
	router.Delete("/node/:id", h.DeleteOne)
}

func (h *nodeHandler) CreateOne(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *nodeHandler) DeleteOne(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *nodeHandler) UpdateOne(c *fiber.Ctx) error {
	panic("unimplemented")
}
