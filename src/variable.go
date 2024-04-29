package engine

func (g *Structure) Variable() {
	for i := 0; i < len(g.artists); i++ {
		g.artistsName = append(g.artistsName, g.artists[i].Name)
		g.dateCreation = append(g.dateCreation, g.artists[i].CreationDate)
		g.artistsMembers = append(g.artistsMembers, g.artists[i].Members)
		g.artistsId = append(g.artistsId, g.artists[i].Id)
		g.artistsPosters = append(g.artistsPosters, g.artists[i].Image)
		g.firstAlbum = append(g.firstAlbum, g.artists[i].FirstAlbum)
		g.locations = append(g.locations, g.artists[i].Locations.Locations)
		g.dates = append(g.dates, g.artists[i].ConcertDates.Dates)
	}
}