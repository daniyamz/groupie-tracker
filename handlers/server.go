package handlers

import (
	"encoding/json"
	"fmt"
	datainfor "groupie-tracker/data"
	"html/template"
	"net/http"
	"os"
)

// Home page handler
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	//validating request path
	if r.URL.Path != "/" {
		http.Error(w, "PATH NOT FOUND", http.StatusNotFound)
		return
	}
	// validating method request
	if r.Method != http.MethodGet {
		http.Error(w, "METHOD NOT FOUND", http.StatusNotFound)
		return
	}
	//set index.html as template and rendering template
	templ, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}
	templ.Execute(w, nil)
}
func jsonOpen(file string, date datainfor.Datas) datainfor.Datas {
	jsonfile, err := os.Open("./jsoncontents/" + file)
	if err != nil {
		fmt.Printf("error occured: %v", err)
	}
	defer jsonfile.Close()

	var jsondata datainfor.Datas
	decoder := json.NewDecoder(jsonfile)

	err = decoder.Decode(&jsondata)

	return jsondata
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	mainp, err := template.ParseFiles("templates/main.html")
	if err != nil {
		http.Error(w, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}
	if r.URL.Path != "/mainpage" {
		http.Error(w, "PATH NOT FOUND", http.StatusNotFound)
		return
	}
	var data datainfor.Datas

	data = jsonOpen("artists.json", data)
	data = jsonOpen("locations.json", data)
	mainp.Execute(w, data)
}
