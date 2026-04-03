package handlers

import (
	"fmt"
	"net/http"

	"github.com/xncs120/goxpress/config"

	"github.com/labstack/echo/v5"
)

func HomeDoc(c *echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func Index(c *echo.Context) error {
	return c.JSON(http.StatusOK, fmt.Sprintf("%s api is running", config.App.Name))
}
