package handler

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type Mapping struct {
	Method  string
	Path    string
	Handler func(c *fiber.Ctx) error
}

type Handler interface {
	Table() []Mapping
}

func AsHandler(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Handler)),
		fx.ResultTags(`group:"handlers"`),
	)
}
