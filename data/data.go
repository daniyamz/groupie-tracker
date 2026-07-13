package groupie

type Artists struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	FirstAlbum   string   `json:"FirstAlbum"`
	Image        string   `json:"image"`
	Members      []string `json:"member"`
	CreationDate int      `json:"creationdate"`
	ConcertDate  string   `json:"concertdate"`
	Relationship string   `json:"releationship"`
}

type Locations struct {
	LastConcertLocations string `json:"lastconcertlocation"`
}
