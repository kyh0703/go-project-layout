package perf

import (
	"context"
	"fmt"
	"net"
	"net/http"

	_ "net/http/pprof"

	"go.uber.org/fx"
)

func NewPerfServer(lc fx.Lifecycle) *http.Server {
	srv := &http.Server{Addr: ":6060"}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			fmt.Println("Start Perf server at", srv.Addr)
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}
