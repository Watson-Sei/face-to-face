package main

import (
	"github.com/Watson-Sei/face-to-face/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)
	app.Post("/facts", handlers.CreateFact)
}
