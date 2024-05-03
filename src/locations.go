package engine

import (
	"strings"
)

type CountriesStruct struct {
	Name string
	Cities []CitiesStruct
}

type CitiesStruct struct {
	CityName string
	Artists []ArtistsStruct
}


func (g *Structure) getPlaces() {
	for i := 0; i < len(g.artists); i++ {
		for j := 0; j < len(g.artists[i].Locations.Locations); j++ {
			temp := strings.Split(g.artists[i].Locations.Locations[j], "-")
			country := CountriesStruct{Name: temp[1],}
			city := CitiesStruct{CityName: temp[0],}
			if !containsContry(g.countries, temp[1]) {g.countries = append(g.countries, country)}
			indice := getIndiceCountryByName(g.countries, country.Name)
			if !containsCity(g.countries[indice].Cities, city) {g.countries[indice].Cities = append(g.countries[indice].Cities, city)}
		}
	}
}

func (g *Structure) setArtistsByCountry() {
	for i := 0; i < len(g.artists); i++ {
		for j := 0; j < len(g.artists[i].Locations.Locations); j++ {
			temp := strings.Split(g.artists[i].Locations.Locations[j], "-")
			indice := getIndiceCountryByName(g.countries ,temp[1])
			for k := 0; k < len(g.countries[indice].Cities); k++ {
				if g.countries[indice].Cities[k].CityName == temp[0] {
					g.countries[indice].Cities[k].Artists = append(g.countries[indice].Cities[k].Artists, g.artists[i])
				}
			}
		}
	}
}

func (g *Structure) fullMaj() {
	for i := 0; i < len(g.countries); i++ {
		g.countries[i].Name = strings.Title(strings.Replace(g.countries[i].Name, "_", "-", -1))
		for j := 0; j < len(g.countries[i].Cities); j++ {
			g.countries[i].Cities[j].CityName = strings.Title(strings.Replace(g.countries[i].Cities[j].CityName, "_", "-", -1))
		}
	}
}




func containsContry(arr []CountriesStruct, element string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i].Name == element {
			return true
		}
	}
	return false
}

func containsCity(arr []CitiesStruct, element CitiesStruct) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i].CityName == element.CityName {
			return true
		}
	}
	return false
}

func getIndiceCountryByName(arr []CountriesStruct, name string) int{
	for i := 0; i < len(arr); i++ {
		if arr[i].Name == name {
			return i
		}
	}
	return 0 
}

