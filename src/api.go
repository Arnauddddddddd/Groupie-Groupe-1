package engine

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type ApiStruct struct {
	Artists   ArtistsStruct   `json:"artists"`
	Locations LocationsStruct `json:"locations"`
	Dates     DatesStruct 	  `json:"dates"`
	Relation  RelationStruct  `json:"relation"`
}

type ArtistsStruct struct {
	Id 			 int 		`json:"id"`
	Image        string   	`json:"image"`
	Name         string     `json:"name"`
	Members      []string	`json:"members"`
	CreationDate int		`json:"creationDate"`
	FirstAlbum   string		`json:"firstAlbum"`
	Locations 	 string		`json:"locations"`
	ConcertDates string		`json:"concertDates"`
	Relations    string		`json:"relations"`
}

type LocationsStruct struct {
	Id 			 int		`json:"id"`
	Locations    []string	`json:"locations"`
	Dates 		 string		`json:"dates"`
}

type DatesStruct struct {
	Id 			 int		`json:"id"`
	Dates 		 []string	`json:"dates"`
}

type RelationStruct struct {
	Id			   int		`json:"id"`
	DatesLocations string	`json:"datesLocations"`
}

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
	fmt.Println(len(list))
	for i := 1; i <= len(list); i++ {
		g.getArtists(i)
		g.getLocation(i)
	}

	fmt.Println(g.api[0].Locations)

}

func (g *Structure) getArtists(i int) {
	var apiStruct ApiStruct
	var artisteStruct ArtistsStruct 
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + strconv.Itoa(i))
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {log.Fatal(err)}
		json.Unmarshal(bodyBytes, &artisteStruct)
	}
	g.api = append(g.api, apiStruct)
	g.api[i-1].Artists = artisteStruct
}

func (g *Structure) getLocation(i int) {
	var locationStruct LocationsStruct 
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + strconv.Itoa(i))
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {log.Fatal(err)}
		json.Unmarshal(bodyBytes, &locationStruct)
	}
	g.api[i-1].Locations = locationStruct
}



