package handler

import (
	"go.uber.org/fx"
)

var HandlerModule = fx.Module(
	"handler",
	fx.Provide(
		AsRoute(NewAuthHandler),
		AsRoute(NewEdgeHandler),
		AsRoute(NewNodeHandler),
		AsRoute(NewSubFlowHandler),
		AsRoute(NewUserHandler),
	),
)
