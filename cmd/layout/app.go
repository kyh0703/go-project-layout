package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type app struct {
	fiber *fiber.App
}

func NewApp() *app {
	return &app{
		fiber: fiber.New(),
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
	a.fiber.Use(cors.New())
	a.fiber.Use(a.setupLogger())
	a.fiber.Use(pprof.New())
	a.fiber.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	a.fiber.Listen(":3000")
	return nil
}

func (a *app) Stop(ctx context.Context) error {
	a.fiber.Shutdown()
	return nil
}
