package middleware

import (
	"net/http"
	"strings"

	"github.com/Watson-Sei/face-to-face/utils"
	"github.com/labstack/echo/v4"
)

// 検証
func JwtMiddleware(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// JWTを取得する処理
			authorization := c.Request().Header.Get("Authorization")
			token := strings.TrimPrefix(authorization, "Bearer ")

			// デコードする
			claims, err := utils.DecodeJWT(token)
			if err != nil {
				return err
			}

			// 検証
			if err := utils.VerifyJWT(claims, role); err != nil {
				return c.NoContent(http.StatusUnauthorized)
			}

			c.Set("jwt", claims)

			return next(c)
		}
	}
}
