package engine

type Structure struct {
	api []ApiStruct
	numberOfGroup int
}

func (g *Structure) Run() {
	g.Init()
	g.Api()
	g.Server()
}
