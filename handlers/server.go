package handlers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/config"
	"html/template"
	"net/http"
)

type Datas struct {
	Artists   []Artist   `json:"artists"`
	Locations []Location `json:"locations"`
	Dates     []Date     `json:"dates"`
	Relations []Relation `json:"relations"`
}

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// function for the landing page
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	// template to parse index.html
	temp, err := template.ParseFiles("templates/index.html")
	if r.URL.Path != "/" {
		http.Error(w, "PATH NOT FOUND", http.StatusNotFound)
		return
	}
	//validating template
	if err != nil {
		http.Error(w, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}
	temp.Execute(w, nil)
}

// function to get data fetch json data
func getJsonData(str string, data Datas) (Datas, error) {
	resp, err := http.Get(str)
	if err != nil {
		return data, fmt.Errorf("error: %v", err)
	}
	defer resp.Body.Close()

	var jdata Datas

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, fmt.Errorf("error occured: %v", err)
	}
	return jdata, nil
}

// mainPageHandler for main.html
func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	maint, err := template.ParseFiles("templates/main.html")
	if r.URL.Path != "/artists" {
		http.Error(w, "PATH NOT FOUND", http.StatusNotFound)
		return
	}
	var adata Datas

	adata, _ = getJsonData(config.Api+"/artists", adata)
	adata, _ = getJsonData(config.Api+"/locations", adata)

	if err != nil {
		http.Error(w, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}
	maint.Execute(w, adata)
}
