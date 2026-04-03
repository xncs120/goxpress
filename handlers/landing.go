package handlers

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func HomeDoc(c *echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func Index(c *echo.Context) error {
	return c.JSON(http.StatusOK, "goxpress api is running")
}
