package handlers

import (
	"net/http"
	"text/template"
)

func LandingPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "PATH NOT FOUND", http.StatusNotFound)
		return
	}
	templ, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}
	templ.Execute(w, nil)
}
