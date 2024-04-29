package engine

type Structure struct {
	artists []ArtistsStruct
	numberOfGroup int
	artistsName []string
	artistsMembers [][]string
	dateCreation []int

}

func (g *Structure) Run() {
	g.Init()
	g.Api()
	g.Variable()
	g.Server()
}
