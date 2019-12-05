package web

import (
	"net/http"
	"text/template"
)

// renderHTML Render HTML web page
func renderHTML(w http.ResponseWriter, tmpl *template.Template) error {
	err := tmpl.Execute(w, nil)
	if err != nil {
		return err
	}

	return nil
}

// HomeView Load HTML home page
func HomeView(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../../web/app/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = renderHTML(w, tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
