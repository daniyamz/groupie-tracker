package datainfor

type Datas struct {
	ArtistsData  []Artists
	LocationData []Location
	DatesDatas   []Dates
}

type Artists struct {
	Id          int      `json:"id"`
	Image       string   `json:"image"`
	Name        string   `json:"names"`
	Members     []string `json:"members"`
	FirstAlbum  string   `json:"dateOfAlbum"`
	ConcertDate string   `json:"concertdate"`
	Relation    string   `json:"relation"`
	CreaDate    int      `json:"creadate"`
}

type Location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locarion"`
	Dates     string   `json:"dates"`
}
type Dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}
