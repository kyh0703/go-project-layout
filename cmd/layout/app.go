package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type app struct {
	fiber   *fiber.App
	perfSvr *PerfServer
}

func NewApp(perfSvr *PerfServer) *app {
	return &app{
		perfSvr: perfSvr,
	}
}

func (a *app) setupLogger() fiber.Handler {
	return logger.New(logger.Config{
		Format:     "${time} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Seoul",
	})
}

func (a *app) Run(ctx context.Context) error {
	a.fiber = fiber.New()
	a.fiber.Use(cors.New())
	a.fiber.Use(a.setupLogger())
	return nil
}

func (a *app) Stop(ctx context.Context) error {
	a.perfSvr.Shutdown(ctx)
	return nil
}
