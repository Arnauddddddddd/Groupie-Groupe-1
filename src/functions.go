package engine

import (
	"strconv"
	"strings"
)

// This function reverses the current list of Artists on the site
func (g *Structure) reverse() {
	for i, j := 0, len(g.artistsTemp)-1; i < j; i, j = i+1, j-1 {
		g.artistsTemp[i], g.artistsTemp[j] = g.artistsTemp[j], g.artistsTemp[i]
	}
}

// This function sorts artists based on string in settings --> call function "sortType"
func (g *Structure) sortArtists(action string) {
	for i := 0; i < len(g.artistsTemp); i++ {
		for j := 0; j < len(g.artistsTemp)-1; j++ {
			if sortType(g.artistsTemp, i, j, action) {
				g.artistsTemp[i], g.artistsTemp[j] = g.artistsTemp[j], g.artistsTemp[i]
			}
		}
	}
}

// this function allows you to get group names and/or people that contain the string as a parameter
func (g *Structure) searchGroup(str string) {
	g.artistsTemp = []ArtistsStruct{}
	for i := 0; i < len(g.artists); i++ {
		if strings.Contains(strings.ToLower(g.artists[i].Name), strings.ToLower(str)) {
			g.artistsTemp = append(g.artistsTemp, g.artists[i])
		}
		for j := 0; j < len(g.artists[i].Members); j++ {
			if strings.Contains(strings.ToLower(g.artists[i].Members[j]), strings.ToLower(str)) && !containsArtist(g.artistsTemp, g.artists[i]) {
				g.artistsTemp = append(g.artistsTemp, g.artists[i])
			}
		}
	}
}

// this function allows you to keep only the artists who created their group between date 1 and 2
func (g *Structure) filterDateCreation(date int, date2 int) {
	artistsTemp2 := []ArtistsStruct{}
	for _, artist := range g.artistsTemp {
		if artist.CreationDate >= date && artist.CreationDate <= date2 {
			artistsTemp2 = append(artistsTemp2, artist)
		}
	}
	g.artistsTemp = artistsTemp2
}

// this function allows you to keep only the artists who make their first album between date 1 and 2
func (g *Structure) filterFirstAlbum(date int, date2 int) {
	artistsTemp2 := []ArtistsStruct{}
	for _, artist := range g.artistsTemp {
		dateArtist, _ := strconv.Atoi(artist.FirstAlbum[6:])
		if dateArtist >= date && dateArtist <= date2 {
			artistsTemp2 = append(artistsTemp2, artist)
		}
	}
	g.artistsTemp = artistsTemp2
}

// this function allows you to only keep artists whose group contains between "nb" and "nb2" people
func (g *Structure) numberMembers(nb int, nb2 int) {
	artistsTemp2 := []ArtistsStruct{}
	for _, artist := range g.artistsTemp {
		if len(artist.Members) >= nb && len(artist.Members) <= nb2 {
			artistsTemp2 = append(artistsTemp2, artist)
		}
	}
	g.artistsTemp = artistsTemp2
}

// reset the list of artists
func (g *Structure) clearArtists() {
	g.artistsTemp = g.artists
}

// this function checks if the two ArtistStructs are the same. It allows you to avoid duplicates in the searchGroup function
func containsArtist(elems []ArtistsStruct, v ArtistsStruct) bool {
	for i := 0; i < len(elems); i++ {
		if elems[i].Id == v.Id {
			return true
		}
	}
	return false
}

// this function checks the conditions for the sortArtists function
func sortType(list []ArtistsStruct, i int, j int, action string) bool {
	switch action {
	case "Year":
		return list[i].CreationDate < list[j].CreationDate
	case "Alphabet":
		return list[i].Name < list[j].Name
	}
	return false
}
