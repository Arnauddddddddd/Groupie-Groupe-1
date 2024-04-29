package engine

func (g *Structure) Variable() {
	for i := 0; i < len(g.artists); i++ {
		g.artistsName = append(g.artistsName, g.artists[i].Name)
		g.dateCreation = append(g.dateCreation, g.artists[i].CreationDate)
		g.artistsMembers = append(g.artistsMembers, g.artists[i].Members)
	}
}