package main

import (
	"fmt"
	"groupie-tracker/handlers"
	"net/http"
)

func main() {
	fmt.Println("Starting server:")
	apimux := http.NewServeMux()
	apimux.HandleFunc("/", handlers.HomeHandler)
	apimux.HandleFunc("/artists", handlers.ArtistHandler)
	http.ListenAndServe(":8090", apimux)
}
