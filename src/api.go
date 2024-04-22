package engine

import (
	"encoding/json"
	"fmt"
)

type ApiStructure struct {
	Artists string
	Locations string
	Dates string
	Relation string
  }

func (g *Structure) Api() {
	data := `{"artists":"https://groupietrackers.herokuapp.com/api/artists","locations":"https://groupietrackers.herokuapp.com/api/locations","dates":"https://groupietrackers.herokuapp.com/api/dates","relation":"https://groupietrackers.herokuapp.com/api/relation"}`
	var apiStructure ApiStructure	
	json.Unmarshal([]byte(data), &apiStructure)
	fmt.Printf("Artists: %s, Locations: %s, Artists: %s, Locations: %s", apiStructure.Artists, apiStructure.Locations, apiStructure.Dates, apiStructure.Relation)
	//Species: pigeon, Description: likes to perch on rocks
}