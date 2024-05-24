package engine

import (
	"strconv"
	"strings"
)

func (g *Structure) reverse() {
	g.artistsTemp = g.artists
	for i, j := 0, len(g.artistsTemp)-1; i < j; i, j = i+1, j-1 {
		g.artistsTemp[i], g.artistsTemp[j] = g.artistsTemp[j], g.artistsTemp[i]
	}
}


func (g *Structure) sortArtists(action string) {
	g.artistsTemp = g.artists
	for i := 0; i < len(g.artistsTemp); i++ {
		for j := 0; j < len(g.artistsTemp)-1; j++ {
			if sortType(g.artistsTemp, i , j , action) {
				g.artistsTemp[i], g.artistsTemp[j] = g.artistsTemp[j], g.artistsTemp[i]
			}
		}
	}
}


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

func (g *Structure) filterDateCreation(date int) {
	g.artistsTemp = []ArtistsStruct{} 
	for _, artist := range g.artists {
		if (artist.CreationDate == date) {
			g.artistsTemp = append(g.artistsTemp, artist)
		}
	}
}

func (g *Structure) filterFirstAlbum(date int) {
	g.artistsTemp = []ArtistsStruct{} 
	for _, artist := range g.artists {
		dateArtist, _ := strconv.Atoi(artist.FirstAlbum[6:])
		if (dateArtist == date) {
			g.artistsTemp = append(g.artistsTemp, artist)
		}
	}
}

func (g *Structure) numberMembers(nb int) {
	g.artistsTemp = []ArtistsStruct{} 
	for _, artist := range g.artists {
		if (len(artist.Members) == nb) {
			g.artistsTemp = append(g.artistsTemp, artist)
		}
	}
}



func (g *Structure) clearArtists() {
	g.artistsTemp = g.artists
}

func containsArtist(elems []ArtistsStruct, v ArtistsStruct) bool {
	for i := 0; i < len(elems); i++ {
		if elems[i].Id == v.Id {
			return true
		}
	}
    return false
}

func sortType(list []ArtistsStruct, i int, j int, action string) bool {
	switch action {
	case "Year":
		return list[i].CreationDate < list[j].CreationDate
		//return list[i].FirstAlbum[len(list[i].FirstAlbum)-4:] < list[j].FirstAlbum[len(list[j].FirstAlbum)-4:]
	case "Alphabet":
		return list[i].Name < list[j].Name
	}
	return false
}

