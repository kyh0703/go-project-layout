package handler

import "github.com/gofiber/fiber/v2"

type SubFlowHandler interface {
	Handler
}

type subFlowHandler struct{}

func NewSubFlowHandler() SubFlowHandler {
	return &subFlowHandler{}
}

func (h *subFlowHandler) Table() []Mapping {
	return []Mapping{
		{Method: fiber.MethodPost, Path: "/subflow", Handler: h.CreateOne},
		{Method: fiber.MethodPut, Path: "/subflow/:id", Handler: h.UpdateOne},
		{Method: fiber.MethodDelete, Path: "/subflow/:id", Handler: h.DeleteOne},
		{Method: fiber.MethodPost, Path: "/subflow/:id/capture", Handler: h.Capture},
		{Method: fiber.MethodPost, Path: "/subflow/:id/undo", Handler: h.Undo},
		{Method: fiber.MethodPost, Path: "/subflow/:id/redo", Handler: h.Redo},
	}
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
