package engine

func (g *Structure) Init() {
	g.artistsTemp = []ArtistsStruct{}
	g.artists = []ArtistsStruct{}
	g.numberOfGroup = 0
	g.action = ""
}
