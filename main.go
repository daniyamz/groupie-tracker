package main

import (
	"fmt"
	"groupie-tracker/handlers"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting server:")
	apimux := http.NewServeMux()
	apimux.HandleFunc("/", handlers.HomeHandler)
	apimux.HandleFunc("/artists", handlers.ArtistHandler)
	err := http.ListenAndServe(":8090", apimux)
	if err != nil {
		log.Printf("Server fails to start: %v", err)
	}
}
