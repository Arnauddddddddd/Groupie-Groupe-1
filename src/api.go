package engine

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)



type ArtistsStruct struct {
	Id 			 int 				`json:"id"`
	Image        string   			`json:"image"`
	Name         string    			`json:"name"`
	Members      []string			`json:"members"`
	CreationDate int				`json:"creationDate"`
	FirstAlbum   string				`json:"firstAlbum"`
	Locations 	 LocationsStruct	`json:"locations"`
	ConcertDates DatesStruct		`json:"concertDates"`
	Relations    RelationStruct		`json:"relations"`
	Dates       [][]string

}

type LocationsStruct struct {
	Id 			 int				`json:"id"`
	Locations    []string			`json:"locations"`
}

type DatesStruct struct {
	Id 			 int				`json:"id"`
	Dates 		 []string			`json:"dates"`
}

type RelationStruct struct {
	Id			   int						`json:"id"`
	DatesLocations map[string][]string		`json:"datesLocations"`
}


// This function is called at startup, it calls each json data fetch function.
// The first extraction allows you to know the number of artists in the API
func (g *Structure) Api() {
	var list []string 
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {log.Fatal(err)}
		json.Unmarshal(bodyBytes, &list)
	}
	for i := 1; i <= len(list); i++ {
		g.getArtists(i)
		g.getLocation(i)
		g.getDates(i)
		g.getRelations(i)
	}
	g.artistsTemp = g.artists
	g.setLocationStruct()
}


// This function allows you to call all the functions to initialize the list of LocationStructs
func (g *Structure) setLocationStruct() {
	g.getPlaces()
	g.setArtistsByCountry()
	g.dateForCity()
	g.fullMaj()
	g.sortCountries()
	g.countriesTemp = g.countries
}


// This function extracts the artists from the API and puts them in a list of ArtistStruct.
func (g *Structure) getArtists(i int) {
	var artisteStruct ArtistsStruct 
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + strconv.Itoa(i))
	if err != nil {log.Fatal(err)}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {log.Fatal(err)}
		json.Unmarshal(bodyBytes, &artisteStruct)
	}
	g.artists = append(g.artists, artisteStruct)
	g.artists[i-1] = artisteStruct
}

// This function extracts locations from the API and places them in the artistStruct corresponding to its position in the AristStruct list
func (g *Structure) getLocation(i int) {
	var locationStruct LocationsStruct 
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + strconv.Itoa(i))
	if err != nil {log.Fatal(err)}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {log.Fatal(err)}
		json.Unmarshal(bodyBytes, &locationStruct)
	}
	g.artists[i-1].Locations = locationStruct
}

// This function extracts dates from the API and places them in the artistStruct corresponding to its position in the AristStruct list
func (g *Structure) getDates(i int) {
	var datesStruct DatesStruct 
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/dates/" + strconv.Itoa(i))
	if err != nil {log.Fatal(err)}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {log.Fatal(err)}
		json.Unmarshal(bodyBytes, &datesStruct)
	}
	g.artists[i-1].ConcertDates = datesStruct
}


// This function extracts relation from the API and places them in the artistStruct corresponding to its position in the AristStruct list
func (g *Structure) getRelations(i int) {
	var relationStruct RelationStruct 
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + strconv.Itoa(i))
	if err != nil {log.Fatal(err)}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {log.Fatal(err)}
		json.Unmarshal(bodyBytes, &relationStruct)
	}
	g.artists[i-1].Relations = relationStruct
}





