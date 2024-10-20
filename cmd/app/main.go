package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/cors"
)

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		os.Exit(1)
	}
}

func run(ctx context.Context, _ io.Writer, _ []string) error {
	ctx, stop := signal.NotifyContext(
		ctx,
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	mux := http.NewServeMux()
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", handler)

	<-ctx.Done()
	_, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return nil
}
