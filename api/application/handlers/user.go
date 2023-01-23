package handlers

import (
	"net/http"

	"github.com/Watson-Sei/face-to-face/application/usecases"
	"github.com/Watson-Sei/face-to-face/domain/models"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserUsecase usecases.UserUsecase
}

func NewUserHandler(uu usecases.UserUsecase) *UserHandler {
	return &UserHandler{UserUsecase: uu}
}

func (uu *UserHandler) GetUsers(c echo.Context) error {
	users, err := uu.UserUsecase.GetUsers()
	if err != nil {
		return err
	}
	return c.JSON(200, users)
}

func (uu *UserHandler) GetUser(c echo.Context) error {
	user, err := uu.UserUsecase.GetUser(c.Param("id"))
	if err != nil {
		return err
	}
	return c.JSON(200, user)
}

func (uu *UserHandler) CreateUser(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	user, err := uu.UserUsecase.CreateUser(user)
	if err != nil {
		return err
	}
	return c.JSON(200, user)
}

func (uu *UserHandler) DeleteUser(c echo.Context) error {
	err := uu.UserUsecase.DeleteUser(c.Param("id"))
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
