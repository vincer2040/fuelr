package render

import (
	"io"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func New() *Template {
	funcMap := template.FuncMap{
		"FormatDate": func(t time.Time) string {
			return t.Format("01/02/2006")
		},
	}
	tmpl := template.New("").Funcs(funcMap)
	tmpl = template.Must(tmpl.ParseGlob("public/*.html"))
	return &Template{
		templates: tmpl,
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, e echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
