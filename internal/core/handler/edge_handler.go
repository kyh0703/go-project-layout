package handler

import "github.com/gofiber/fiber/v2"

type edgeHandler struct{}

func NewEdgeHandler() Handler {
	return &edgeHandler{}
}

func (h *edgeHandler) Table() []Mapping {
	return []Mapping{
		{Method: fiber.MethodPost, Path: "/edge", Handler: h.CreateOne},
		{Method: fiber.MethodPut, Path: "/edge/:id", Handler: h.UpdateOne},
		{Method: fiber.MethodDelete, Path: "/edge/:id", Handler: h.DeleteOne},
	}
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
