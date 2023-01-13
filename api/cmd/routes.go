package main

import (
	"github.com/Watson-Sei/face-to-face/handlers"
	"github.com/Watson-Sei/face-to-face/middleware"
	"github.com/labstack/echo/v4"
)

func setupRoutes(app *echo.Echo) {
	// guest level routes
	app.GET("/api/guest/check", handlers.Veirfy, middleware.JwtMiddleware("guest"))
	app.GET("/api/staff/check", handlers.Veirfy, middleware.JwtMiddleware("staff"))
	app.GET("/api/owner/check", handlers.Veirfy, middleware.JwtMiddleware("owner"))

	// generate token routes
	auth := app.Group("/api/auth")
	auth.POST("/token", handlers.GetToken)

	// generate token routes
	app.GET("/rtc/:channelName/:role/:tokentype/:uid", handlers.GetRtcToken)
	app.GET("/rtm/:uid", handlers.GetRtmToken)
	app.GET("/rte/:channelName/:role/:tokentype/:uid", handlers.GetBothTokens)

	// guest user routes
	app.GET("/users", handlers.GetUsers)
	app.GET("/user/:id", handlers.GetUser)
	app.POST("/user", handlers.CreateUser)
	app.DELETE("/user/:id", handlers.DeleteUser)
}
