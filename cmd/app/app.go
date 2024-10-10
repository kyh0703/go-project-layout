package main

import "github.com/gofiber/fiber/v2"

type App struct {
	app *fiber.App
}

func New() *App {
	return &App{
		app: fiber.New(),
	}
}
