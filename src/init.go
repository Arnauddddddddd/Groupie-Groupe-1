package engine


func (g *Structure) Init() {
	g.artists = []ArtistsStruct{}
	g.numberOfGroup = 0

	g.artistsId = []int{}
	g.artistsName = []string{}
	g.artistsMembers = [][]string{}
	g.dateCreation = []int{}
	g.artistsPosters = []string{}
	g.firstAlbum = []string{}
	g.locations = [][]string{}
	g.dates = [][]string{}
}