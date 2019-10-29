package main

import (
	"net/http"
	"text/template"
)

type Page struct {
	name string
}

var tmpls = template.Must(template.ParseFiles("page.html"))

//TemplateInit used every where to initialize HTML templates
func TemplateInit(w http.ResponseWriter, templateFile string, templateData Page) {
	if err := tmpls.ExecuteTemplate(w, templateFile, templateData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	page := Page{name: "ExampleVideo"}
	TemplateInit(w, "page.html", page)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9090", nil)
}
