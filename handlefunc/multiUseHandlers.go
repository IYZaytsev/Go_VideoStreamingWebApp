package handlefunc

import (
	"net/http"
	"text/template"
)

//Page struct is used for HTML templates
type Page struct {
	Name string
}

var tmpls = template.Must(template.ParseFiles("page.html"))

//TemplateInit used every where to initialize HTML templates
func TemplateInit(w http.ResponseWriter, templateFile string, templateData Page) {
	if err := tmpls.ExecuteTemplate(w, templateFile, templateData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//Index is the main page of the web app
func Index(w http.ResponseWriter, r *http.Request) {
	page := Page{Name: "assssssssssss"}
	TemplateInit(w, "page.html", page)
}
