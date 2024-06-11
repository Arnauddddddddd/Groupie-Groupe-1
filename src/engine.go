package engine

type Structure struct {
	artists		  []ArtistsStruct
	artistsTemp   []ArtistsStruct
	numberOfGroup int
	action		  string
	countries	  []CountriesStruct
	countriesTemp []CountriesStruct
}

// the main function
func (g *Structure) Run() {
	g.Init()
	g.Api()
	g.Server()
}
