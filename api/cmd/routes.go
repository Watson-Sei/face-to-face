package main

import (
	"github.com/Watson-Sei/face-to-face/handlers"
	"github.com/Watson-Sei/face-to-face/middleware"
	"github.com/labstack/echo/v4"
)

func setupRoutes(app *echo.Echo) {
	// guest level middleware
	app.GET("/api/guest/check", handlers.Veirfy, middleware.JwtMiddleware("guest"))
	app.GET("/api/staff/check", handlers.Veirfy, middleware.JwtMiddleware("staff"))
	app.GET("/api/owner/check", handlers.Veirfy, middleware.JwtMiddleware("owner"))

	auth := app.Group("/api/auth")
	auth.POST("/token", handlers.GetToken)

	app.GET("/users", handlers.GetUsers)
	app.GET("/user/:id", handlers.GetUser)
	app.POST("/user", handlers.CreateUser)
	app.DELETE("/user/:id", handlers.DeleteUser)
}
