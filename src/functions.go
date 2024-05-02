package engine

import "strings"

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
        if strings.Contains(g.artists[i].Name, str) {
            g.artistsTemp = append(g.artistsTemp, g.artists[i])
        }
    }
}


func (g *Structure) clearArtists() {
	g.artistsTemp = g.artists
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


