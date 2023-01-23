package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Watson-Sei/face-to-face/infrastructure/auth"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestJwtMiddlewareInvalidToken(t *testing.T) {
	e := echo.New()
	e.Use(JwtMiddleware("guest"))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// case1: invalid token
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer invalid token")
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestJwtMiddlewareValidToken(t *testing.T) {
	e := echo.New()
	e.Use(JwtMiddleware("guest"))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	token, _ := auth.GenerateJWT("1", "guest")

	// case 2: valid token
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestJwtMiddlewareInvalidRole(t *testing.T) {
	token, _ := auth.GenerateJWT("1", "guest")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	}, JwtMiddleware("admin"))

	// case 3: invalid role
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestJwtMiddlewareValidRole(t *testing.T) {
	token, _ := auth.GenerateJWT("1", "admin")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	// case 4: valid role
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
