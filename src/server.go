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

    btnFiltreDateCreation := r.Form.Get("yearFilterForm")
	if len(btnFiltreDateCreation) > 0 {
		minYear := r.Form.Get("minYear")
		maxYear := r.Form.Get("maxYear")
		nb, _ := strconv.Atoi(minYear)
		nb2, _ := strconv.Atoi(maxYear)
		g.filterDateCreation(nb, nb2)
	}

	filterFirstAlbumForm := r.Form.Get("filterFirstAlbumForm")
	if len(filterFirstAlbumForm) > 0 {
		minFirstAlbumYear := r.Form.Get("minFirstAlbumYear")
		maxFirstAlbumYear := r.Form.Get("maxFirstAlbumYear")
		nb, _ := strconv.Atoi(minFirstAlbumYear)
		nb2, _ := strconv.Atoi(maxFirstAlbumYear)
		g.filterFirstAlbum(nb, nb2)
	}

	btnFiltreNumberMembers:= r.Form.Get("filterNumberMembersForm")
	if len(btnFiltreNumberMembers) > 0 {
		maxNumberMembers := r.Form.Get("maxNumberMembers")
		minNumberMembers := r.Form.Get("minNumberMembers")
		nb, _ := strconv.Atoi(minNumberMembers)
		nb2, _ := strconv.Atoi(maxNumberMembers)
		g.numberMembers(nb, nb2)
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


