package handlers

import (
	"html/template"
	"net/http"
)

type ErrorPage struct {
	ErrorMessage string
	ErrorStatus  int
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}
	temp.Execute(w, nil)
}
