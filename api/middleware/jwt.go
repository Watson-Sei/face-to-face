package middleware

import (
	"strings"

	"github.com/Watson-Sei/face-to-face/infrastructure/auth"
	"github.com/labstack/echo/v4"
)

// 検証
func JwtMiddleware(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			authorization := c.Request().Header.Get("Authorization")
			token := strings.TrimPrefix(authorization, "Bearer ")

			claims, err := auth.ValidateJWT(token, role)
			if err != nil {
				return c.JSON(401, map[string]string{"message": err.Error()})
			}

			c.Set("jwt", claims)

			return next(c)
		}
	}
}
