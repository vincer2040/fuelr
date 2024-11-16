package render

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func New() *Template {
	tmpl := template.New("")
	tmpl = template.Must(tmpl.ParseGlob("public/*.html"))
	return &Template{
		templates: tmpl,
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, e echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
