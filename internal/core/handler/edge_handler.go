package handler

import "github.com/gofiber/fiber/v2"

type EdgeHandler interface {
	Handler
	CreateOne(c *fiber.Ctx) error
	GetOne(c *fiber.Ctx) error
	DeleteOne(c *fiber.Ctx) error
	UpdateOne(c *fiber.Ctx) error
}

type edgeHandler struct{}

func NewEdgeHandler() EdgeHandler {
	return &edgeHandler{}
}

func (h *edgeHandler) Table() []Mapping {
	return []Mapping{
		{Method: fiber.MethodPost, Path: "/edge", Handler: h.CreateOne},
		{Method: fiber.MethodGet, Path: "/edge/:id", Handler: h.GetOne},
		{Method: fiber.MethodPut, Path: "/edge/:id", Handler: h.UpdateOne},
		{Method: fiber.MethodDelete, Path: "/edge/:id", Handler: h.DeleteOne},
	}
}

func (h *edgeHandler) CreateOne(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *edgeHandler) GetOne(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *edgeHandler) DeleteOne(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *edgeHandler) UpdateOne(c *fiber.Ctx) error {
	panic("unimplemented")
}
