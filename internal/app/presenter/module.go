package presenter

import (
	"github.com/kyh0703/go-project-layout/internal/app/presenter/perf"
	"github.com/kyh0703/go-project-layout/internal/app/presenter/web"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"presenter",
	fx.Provide(
		web.NewHTTPServer,
		perf.NewPerfServer,
	),
)
