package engine

type Structure struct {
	artists		  []ArtistsStruct
	artistsTemp   []ArtistsStruct
	numberOfGroup int
	action		  string
}

func (g *Structure) Run() {
	g.Init()
	g.Api()
	g.Server()
}
