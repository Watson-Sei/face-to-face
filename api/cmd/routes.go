package main

import (
	"github.com/Watson-Sei/face-to-face/handlers"
	"github.com/labstack/echo/v4"
)

func setupRoutes(app *echo.Echo) {
	auth := app.Group("/api/auth")
	auth.POST("/token", handlers.GetToken)

	app.GET("/users", handlers.GetUsers)
	app.GET("/user/:id", handlers.GetUser)
	app.POST("/user", handlers.CreateUser)
	app.DELETE("/user/:id", handlers.DeleteUser)
}
