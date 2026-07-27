package handlers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/config"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Location     string   `json:"locations"`
	ConcertDate  string   `json:"concertDates"`
	Relation     string   `json:"relations"`
}

type Location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
}

type LocationIndex struct {
	Index []Location `json:"index"`
}

type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}
type DateIndex struct {
	Index []Date `json:"index"`
}
type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
type RelationIndex struct {
	Index []Relation `json:"index"`
}

func FetchJsondata(url string, target interface{}) error {
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("fail to fetch %s, StatusCode %d", err, resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(target)
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	htmp, err := template.ParseFiles("templates/index.html")
	if r.URL.Path != "/" {
		http.Error(w, "PATH NOT ALLOWED", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "INTERNAL SERVER ERROR: FETCHING ARTISTS DATA", http.StatusInternalServerError)
		return
	}
	htmp.Execute(w, nil)
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	artist, err := template.ParseFiles("templates/artists.html")

	if err != nil {
		log.Printf("Template parsing failed: %v", err)
		http.Error(w, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "METHOD NOT ALLOWED", http.StatusMethodNotAllowed)
		return
	}
	var artdata []Artist
	er := FetchJsondata(config.Api+"/artists", &artdata)
	if er != nil {
		http.Error(w, "INTERNAL SERVER ERROR: FETCHING ARTISTS DATA", http.StatusInternalServerError)
		return
	}
	artist.Execute(w, artdata)
}
