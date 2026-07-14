package main

import (
	"fmt"
	"groupie-tracker/handlers"
	"net/http"
)

func main() {
	fmt.Println("Sever is running on port :8090")

	musreco := http.NewServeMux()

	musreco.HandleFunc("/", handlers.LandingPageHandler)

	err := http.ListenAndServe(":8090", musreco)
	if err != nil {
		fmt.Println("Error in running the server")
	}
}
