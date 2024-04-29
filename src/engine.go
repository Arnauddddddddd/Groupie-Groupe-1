package engine

type Structure struct {
	artists       []ArtistsStruct
	numberOfGroup int
}

func (g *Structure) Run() {
	g.Init()
	g.Api()
	g.Server()
}
