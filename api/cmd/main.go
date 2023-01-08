package main

import (
	"net/http"

	"github.com/Watson-Sei/face-to-face/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.ConnectDb()

	app := echo.New()

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	setupRoutes(app)

	app.Logger.Fatal(app.Start(":3000"))
}
