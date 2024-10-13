package main

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/kyh0703/go-project-layout/configs"
	"github.com/kyh0703/go-project-layout/internal/adaptor/rpc"
)

type App struct {
	fiber      *fiber.App
	perfServer *http.Server
	wg         sync.WaitGroup
}

func NewApp(
	ctx context.Context,
	rpc rpc.Rpc,
) (*App, error) {
	app := new(App)

	configs.Print()

	if err := app.listenServer(ctx); err != nil {
		return nil, err
	}

	return app, nil
}

func (app *App) Close() {
	app.fiber.Shutdown()
	app.perfServer.Close()
	app.wg.Wait()
}

func (app *App) ListenPerfServer() {
	app.perfServer = &http.Server{
		Addr:         ":6060",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.wg.Add(1)
	go func() {
		if err := app.perfServer.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
}

func (app *App) listenServer(
	ctx context.Context,
) error {
	fiber := fiber.New(fiber.Config{})
	fiber.Use(monitor.New())
	fiber.Use(recover.New())
	fiber.Use(logger.New())
	fiber.Use(cors.New())
	app.fiber = fiber

	go func() {
		fiber.Listen(":" + configs.Env.AppPort)
	}()

	return nil
}
