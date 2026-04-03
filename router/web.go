package router

import (
	"html/template"
	"io"

	"github.com/xncs120/goxpress/handlers"
	"github.com/xncs120/goxpress/views"

	"github.com/labstack/echo/v5"
	echoSwagger "github.com/swaggo/echo-swagger/v2"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(c *echo.Context, w io.Writer, name string, data any) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (r *Router) AddWebRoutes() {
	template := &Template{
		templates: template.Must(template.ParseFS(views.HtmlFiles, "*")),
	}
	r.e.Renderer = template

	// embeded index.html
	r.e.GET("/docs", handlers.HomeDoc)
	r.e.GET("/views/docs/*any", func(c *echo.Context) error {
		return c.Blob(200, "text/yaml", views.YamlFiles)
	})
	r.e.GET("/docs/*any", echoSwagger.EchoWrapHandler(func(c *echoSwagger.Config) {
		c.URLs = []string{"/views/docs/root.yaml"}
	}))

	r.e.GET("/", handlers.Index)
	// not embeded frontend client
	// r.e.Static("/assets", "client/dist/assets")
	// r.e.File("/", "client/dist/index.html")
	// r.e.GET("/*", func(c *echo.Context) error {
	// 	return c.File("client/dist/index.html")
	// })
}
