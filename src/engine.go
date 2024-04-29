package engine

type Structure struct {
	artists []ArtistsStruct
	numberOfGroup int

	artistsName []string
	artistsMembers [][]string
	dateCreation []int
	artistsId []int
	artistsPosters []string
	firstAlbum []string
	locations [][]string
	dates [][]string

}

func (g *Structure) Run() {
	g.Init()
	g.Api()
	g.Variable()
	g.Server()
}
