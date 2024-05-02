package engine

func (g *Structure) reverse() {
	for i, j := 0, len(g.artists)-1; i < j; i, j = i+1, j-1 {
		g.artists[i], g.artists[j] = g.artists[j], g.artists[i]
	}
}


func (g *Structure) sortArtists(action string) {
	for i := 0; i < len(g.artists); i++ {
		for j := 0; j < len(g.artists)-1; j++ {
			if sortType(g.artists, i , j , action) {
				g.artists[i], g.artists[j] = g.artists[j], g.artists[i]
			}
		}
	}
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

