package helpers

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func LoadTemplates() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
}

func Render(w http.ResponseWriter, nom string, data interface{}) {
	templates.ExecuteTemplate(w, nom, data)
}
