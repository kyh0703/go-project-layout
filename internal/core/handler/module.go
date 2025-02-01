package handler

import (
	"go.uber.org/fx"
)

var HandlerModule = fx.Module(
	"handler",
	fx.Provide(
		AsHandler(NewAuthHandler),
		AsHandler(NewPostHandler),
		AsHandler(NewUserHandler),
	),
)
