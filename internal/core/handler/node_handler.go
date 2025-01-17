package handler

import "github.com/gofiber/fiber/v2"

type NodeHandler interface {
	Handler
	CreateOne(c *fiber.Ctx) error
	GetOne(c *fiber.Ctx) error
	DeleteOne(c *fiber.Ctx) error
	UpdateOne(c *fiber.Ctx) error
}

type nodeHandler struct{}

func NewNodeHandler() NodeHandler {
	return &nodeHandler{}
}

func (h *nodeHandler) Table() []Mapping {
	return []Mapping{
		{Method: fiber.MethodPost, Path: "/node", Handler: h.CreateOne},
		{Method: fiber.MethodGet, Path: "/node/:id", Handler: h.GetOne},
		{Method: fiber.MethodPut, Path: "/node/:id", Handler: h.UpdateOne},
		{Method: fiber.MethodDelete, Path: "/node/:id", Handler: h.DeleteOne},
	}
}

func (h *nodeHandler) CreateOne(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *nodeHandler) GetOne(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *nodeHandler) DeleteOne(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *nodeHandler) UpdateOne(c *fiber.Ctx) error {
	panic("unimplemented")
}
