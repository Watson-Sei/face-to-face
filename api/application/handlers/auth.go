package handlers

import (
	"github.com/Watson-Sei/face-to-face/application/usecases"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthUsecase usecases.AuthUsecase
}

func NewAuthHandler(au usecases.AuthUsecase) *AuthHandler {
	return &AuthHandler{AuthUsecase: au}
}

type LoginRequest struct {
	Code string `json:"code"`
}

func (au *AuthHandler) Login(c echo.Context) error {
	loginReq := new(LoginRequest)
	if err := c.Bind(loginReq); err != nil {
		return err
	}
	token, err := au.AuthUsecase.Login(loginReq.Code)
	if err != nil {
		return err
	}
	return c.JSON(200, map[string]string{"access_token": token})
}

func (au *AuthHandler) Veirfy(c echo.Context) error {
	return c.NoContent(200)
}
