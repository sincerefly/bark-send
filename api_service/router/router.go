package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func InitRouter(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}