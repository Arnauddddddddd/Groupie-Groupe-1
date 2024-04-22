package engine

type Structure struct {
}

func (g *Structure) Run() {
	g.Init()
	g.Api()
	g.Server()
}
