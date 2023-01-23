package main

import (
	"github.com/Watson-Sei/face-to-face/application/handlers"
	"github.com/Watson-Sei/face-to-face/application/usecases"
	"github.com/Watson-Sei/face-to-face/domain/models"
	"github.com/Watson-Sei/face-to-face/domain/repositories"
	"github.com/Watson-Sei/face-to-face/infrastructure/database/gorm"
	jwt "github.com/Watson-Sei/face-to-face/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := gorm.ConnectDb()
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})

	app := echo.New()

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	userRepo := handlers.NewUserHandler(usecases.NewUserUsecase(repositories.NewUserRepository(db)))
	authRepo := handlers.NewAuthHandler(usecases.NewAuthUsecase(repositories.NewUserRepository(db)))

	app.GET("/users", userRepo.GetUsers)
	app.GET("/users/:id", userRepo.GetUser)
	app.POST("/users", userRepo.CreateUser)
	app.DELETE("/users/:id", userRepo.DeleteUser)

	// auth routes
	app.POST("/api/auth/login", authRepo.Login)
	app.GET("/api/auth/guest/check", authRepo.Veirfy, jwt.JwtMiddleware("guest"))
	app.GET("/api/auth/staff/check", authRepo.Veirfy, jwt.JwtMiddleware("staff"))

	app.Logger.Fatal(app.Start(":3000"))
}
