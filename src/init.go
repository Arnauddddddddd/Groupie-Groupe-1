package engine


func (g *Structure) Init() {
	g.artists = []ArtistsStruct{}
	g.numberOfGroup = 0

	g.artistsName = []string{}
	g.artistsMembers = [][]string{}
	g.dateCreation = []int{}
}