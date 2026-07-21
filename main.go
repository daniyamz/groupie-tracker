package main

import (
	"fmt"
	"groupie-tracker/handlers"
	"log"
	"net/http"
)

func main() {
	//create a new instance of a servemux
	grup := http.NewServeMux()

	//register handlers
	grup.HandleFunc("/", handlers.HomePageHandler)
	grup.HandleFunc("/artists", handlers.MainPageHandler)

	fmt.Println("Server running at port :8090")
	fmt.Println("To shut down server Press CTRL + C")
	err := http.ListenAndServe(":8090", grup)
	if err != nil {
		log.Fatal()
	}
}
