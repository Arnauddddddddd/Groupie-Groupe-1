package engine

// initialization of global variables
func (g *Structure) Init() {
	g.artistsTemp = []ArtistsStruct{}       // variable displayed on this site which contains part of "g.artists"
	g.artists = []ArtistsStruct{}			// contains all the countries which have concerts
	g.numberOfGroup = 0						
	g.action = ""
	g.countries = []CountriesStruct{}		// contains all the countries which have concerts
	g.countriesTemp = []CountriesStruct{}   // variable displayed on this site which contains part of "g.countries"
}
