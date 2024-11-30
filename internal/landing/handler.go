package landing

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func (h *Handler) Register(c echo.Context) error {
	return c.Render(http.StatusOK, "register.html", nil)
}

func (h *Handler) Docs(c echo.Context) error {
	return c.Render(http.StatusOK, "schemas.html", nil)
}
