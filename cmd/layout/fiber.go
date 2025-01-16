package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/kyh0703/layout/internal/core/handler"
)

func setupMiddleware(app *fiber.App) *fiber.App {
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format:     "${time} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Seoul",
	}))
	app.Use(pprof.New())
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	return app
}

func NewFiber(handlers ...handler.Handler) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      "layout",
		ServerHeader: "layout",
		Prefork:      false,
	})

	app.Get("/swagger/*", swagger.HandlerDefault)
	app = setupMiddleware(app)
	return app
}
