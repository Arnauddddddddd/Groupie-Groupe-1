package engine

import (
	"strings"
)

type CountriesStruct struct {
	Name   string
	Cities []CitiesStruct
}

type CitiesStruct struct {
	CityName string
	Artists  []ArtistsStruct
}

// retrieving all cities and countries from ArtistStruct to extract them into CountryStruct
func (g *Structure) getPlaces() {
	for i := 0; i < len(g.artists); i++ {
		for j := 0; j < len(g.artists[i].Locations.Locations); j++ {
			temp := strings.Split(g.artists[i].Locations.Locations[j], "-")
			country := CountriesStruct{Name: temp[1]}
			city := CitiesStruct{CityName: temp[0]}
			if !containsContry(g.countries, temp[1]) {
				g.countries = append(g.countries, country)
			}
			indice := getIndiceCountryByName(g.countries, country.Name)
			if !containsCity(g.countries[indice].Cities, city) {
				g.countries[indice].Cities = append(g.countries[indice].Cities, city)
			}
		}
	}
}

// retrieving all artists to CountryStruct
func (g *Structure) setArtistsByCountry() {
	for i := 0; i < len(g.artists); i++ {
		for j := 0; j < len(g.artists[i].Locations.Locations); j++ {
			temp := strings.Split(g.artists[i].Locations.Locations[j], "-")
			indice := getIndiceCountryByName(g.countries, temp[1])
			for k := 0; k < len(g.countries[indice].Cities); k++ {
				if g.countries[indice].Cities[k].CityName == temp[0] {
					g.countries[indice].Cities[k].Artists = append(g.countries[indice].Cities[k].Artists, g.artists[i])
				}
			}
		}
	}
}

// remove all tickets from cities
func (g *Structure) removeTiret() {
	for i := 0; i < len(g.countries); i++ {
		g.countries[i].Name = strings.Title(strings.Replace(g.countries[i].Name, "_", "-", -1))
		for j := 0; j < len(g.countries[i].Cities); j++ {
			g.countries[i].Cities[j].CityName = strings.Title(strings.Replace(g.countries[i].Cities[j].CityName, "_", "-", -1))
		}
	}
}

// this function add all artist dates based on city
func (g *Structure) dateForCity() {
	for i := 0; i < len(g.countries); i++ {
		for j := 0; j < len(g.countries[i].Cities); j++ {
			for k := 0; k < len(g.countries[i].Cities[j].Artists); k++ {
				city := g.countries[i].Cities[j].Artists[k].Relations.DatesLocations[g.countries[i].Cities[j].CityName+"-"+g.countries[i].Name]
				for m := 0; m < len(city); m++ {
					city[m] = g.changeFormatDate(city[m])
				}
				g.countries[i].Cities[j].Artists[k].Dates = append(g.countries[i].Cities[j].Artists[k].Dates, city)
			}
		}
	}
}

// this function replace date in handwriting
func (g *Structure) changeFormatDate(date string) string {
	date2 := strings.Split(date, "-")
	if len(date2) != 3 {
		return date
	}
	month := ""
	switch date2[1] {
	case "01":
		month = "January"
	case "02":
		month = "Februry"
	case "03":
		month = "March"
	case "04":
		month = "April"
	case "05":
		month = "May"
	case "06":
		month = "June"
	case "07":
		month = "July"
	case "08":
		month = "August"
	case "09":
		month = "September"
	case "10":
		month = "October"
	case "11":
		month = "November"
	case "12":
		month = "December"
	}
	return date2[0] + (", ") + month + " " + date2[2]
}

// this function sorts countries alphabetically 
func (g *Structure) sortCountries() {
	for i := 0; i < len(g.countries); i++ {
		for j := 0; j < len(g.countries)-1; j++ {
			if g.countries[i].Name < g.countries[j].Name {
				g.countries[i], g.countries[j] = g.countries[j], g.countries[i]
			}
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

func getIndiceCountryByName(arr []CountriesStruct, name string) int {
	for i := 0; i < len(arr); i++ {
		if arr[i].Name == name {
			return i
		}
	}
	return 0
}

// this function adds the country that corresponds to the string in the list of current countries
func (g *Structure) addCurrentCountries(country string) {
	g.countriesTemp = []CountriesStruct{}
	for i := 0; i < len(g.countries); i++ {
		if g.countries[i].Name == country {
			g.countriesTemp = append(g.countriesTemp, g.countries[i])
		}
	}
}
