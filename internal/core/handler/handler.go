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
		fx.ResultTags(`group:"watchers"`),
	)
}

func NewHandler(app *fiber.App, handler []Handler) {
	for _, h := range handler {
		for _, m := range h.Table() {
			switch m.Method {
			case fiber.MethodGet:
				app.Get(m.Path, m.Handler)
			case fiber.MethodPost:
				app.Post(m.Path, m.Handler)
			case fiber.MethodPut:
				app.Put(m.Path, m.Handler)
			case fiber.MethodPatch:
				app.Patch(m.Path, m.Handler)
			case fiber.MethodOptions:
				app.Options(m.Path, m.Handler)
			case fiber.MethodDelete:
				app.Delete(m.Path, m.Handler)
			}
		}
	}
}
