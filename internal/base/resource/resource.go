package resource

import (
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"goxpress/assets"
)

func StaticHandler() echo.HandlerFunc {
	subFS, err := fs.Sub(assets.EmbeddedFiles, "statics")
	if err != nil {
		log.Fatalf("Failed to embed static files: %v", err)
	}
	return echo.WrapHandler(http.StripPrefix("/statics/", http.FileServer(http.FS(subFS))))
}

type TemplateRenderer struct {
	templates *template.Template
}

func NewTemplateRenderer() *TemplateRenderer {
	tmpl := template.Must(template.ParseFS(assets.EmbeddedFiles, "templates/*.html", "templates/**/*.html"))
	return &TemplateRenderer{templates: tmpl}
}

func (f *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return f.templates.ExecuteTemplate(w, name, data)
}
