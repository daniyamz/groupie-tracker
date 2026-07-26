package handlers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/config"
	"html/template"
	"net/http"
)

type Datas struct {
	Artist
	LocationsDatas Location
	DatesDatas     Date
	RelationsDatas Relation
}

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
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
func fetchApi(url string, target any) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		return fmt.Errorf("error occured: %v", err)
	}
	return nil
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	artist, err := template.ParseFiles("templates/artists.html")

	//validating the get method
	if r.Method != http.MethodGet {
		http.Error(w, "METHOD NOT ALLOWED", http.StatusMethodNotAllowed)
		return
	}

	//fetch artists data
	var artdat []Artist
	err = fetchApi(config.Api+"/artists", &artist)
	if err != nil {
		http.Error(w, "INTERNAL SERVER ERROR: FAILD TO PARSE DATA", http.StatusInternalServerError)
		return
	}
	var locIdx LocationIndex
	if err := fetchApi(config.Api+"/loctions", &locIdx); err != nil {
		http.Error(w, "FETCH ERROR: LOCATIONS", http.StatusInternalServerError)
		return
	}
	var dateIdx DateIndex
	err = fetchApi(config.Api+"/dates", &dateIdx)
	if err != nil {
		http.Error(w, "FETCH ERROR: DATES", http.StatusInternalServerError)
		return
	}
	var relIdx RelationIndex
	err = fetchApi(config.Api+"/relations", &relIdx)
	if err != nil {
		http.Error(w, "FETCH ERROR: RELATIONS", http.StatusInternalServerError)
		return
	}
	var finalData []Datas

	for i := 0; i < len(artdat); i++ {
		artistdata := Datas{
			Artist: artdat[i],
		}
		for _, l := range locIdx.Index {
			if l.Id == artdat[i].Id {
				artistdata.LocationsDatas = l
				break
			}
		}
		for _, d := range dateIdx.Index {
			if d.Id == artdat[i].Id {
				artistdata.DatesDatas = d
				break
			}
		}
		for _, r := range relIdx.Index {
			if r.Id == artdat[i].Id {
				artistdata.RelationsDatas = r
				break
			}
		}
		finalData = append(finalData, artistdata)
	}
	err = artist.Execute(w, finalData)
	if err != nil {
		http.Error(w, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}
}
