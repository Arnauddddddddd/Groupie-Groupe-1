package engine

type Structure struct {
}

func (g *Structure) Run() {
	for {
		g.Server()
	}
}
