package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"net/http"
	"html/template"
	"time"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	renderer := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer

	delaySeconds := 3	

	e.GET(`/:text`, func(c echo.Context) error {
			param := c.Param("text")

			data := map[string]string{
				"param": param,
			}

			time.Sleep(time.Duration(delaySeconds) * time.Second)

			return c.Render(http.StatusOK, "index.html", data)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
