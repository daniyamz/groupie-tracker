package main

import (
	"fmt"
	"groupie-tracker/handlers"
	"net/http"
)

func main() {
	fmt.Println("Starting Server...")

	fmt.Println("On htt://localhost:8090")

	fmt.Println("To shut down the server press CTRL + C")

	musreco := http.NewServeMux()

	musreco.HandleFunc("/", handlers.HomePageHandler)

	err := http.ListenAndServe(":8090", musreco)
	if err != nil {
		fmt.Println("Error in running the server")
	}
}
