package engine

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// find server on ubuntu : sudo ss -tulpn

type WebStruct struct {
	Artists []ArtistsStruct
}

type WebStruct2 struct {
	Countries 	 []CountriesStruct
	CountriesTemp []CountriesStruct
}

func (g *Structure) Server() {
	http.HandleFunc("/", g.index)
	http.HandleFunc("/index.html", g.index)
	http.HandleFunc("/locations.html", g.locations)
	http.Handle("/videos/", http.StripPrefix("/videos/", http.FileServer(http.Dir("videos"))))
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("scripts"))))

	fmt.Println("\nhttp://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func (g *Structure) index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("pages/index.html"))
	r.ParseForm()
	action := r.Form.Get("action")
	g.action = action
	if len(g.action) > 0 {
		switch g.action {
		case "sortYear":
			g.sortArtists("Year")
		case "sortAlphabetic":
			g.sortArtists("Alphabet")
		case "reverse":
			g.reverse()
		case "clear":
			g.clearArtists()
		default:
			fmt.Println(g.action)
		}
	}

	search := r.Form.Get("search")
	if len(search) > 0 {
		g.searchGroup(search)
	}

    btnFiltreDateCreation := r.Form.Get("filterDateCreation")
	if len(btnFiltreDateCreation) > 0 {
		date, _ := strconv.Atoi(btnFiltreDateCreation)
		g.filterDateCreation(date)
	}

	btnFiltreFirstAlbum := r.Form.Get("filterFirstAlbum")
	if len(btnFiltreFirstAlbum) > 0 {
		date, _ := strconv.Atoi(btnFiltreFirstAlbum)
		g.filterFirstAlbum(date)
	}

	btnFiltreNumberMembers:= r.Form.Get("filterNumberMembers")
	if len(btnFiltreNumberMembers) > 0 {
		nb, _ := strconv.Atoi(btnFiltreNumberMembers)
		g.numberMembers(nb)
	}

	web := WebStruct{Artists: g.artistsTemp}
	err := tmpl.Execute(w, web)
	if err != nil {
		return
	}
}

func (g *Structure) locations(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	tmpl := template.Must(template.ParseFiles("pages/locations.html"))

	showMovies := r.Form.Get("showMovies")
	if len(showMovies) > 0 {
		if showMovies == "allMovies" {
			g.countriesTemp = g.countries
		} else {
			g.addOrRemove(showMovies)
		}
	}
	
	web2 := WebStruct2{Countries: g.countries, CountriesTemp: g.countriesTemp}
	err := tmpl.Execute(w, web2)
	if err != nil {return}
}


