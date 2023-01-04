package main

import (

	"github.com/Watson-Sei/face-to-face/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.ConnectDb()

	app := echo.New()

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	setupRoutes(app)

	app.Logger.Fatal(app.Start(":3000"))
}
