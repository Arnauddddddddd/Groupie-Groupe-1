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

    filter := r.Form.Get("filter")
	if len(filter) > 0 {
		minYear := r.Form.Get("minYear")
		maxYear := r.Form.Get("maxYear")
		if len(minYear) > 0 && len(maxYear) > 0 {
			nb, _ := strconv.Atoi(minYear)
			nb2, _ := strconv.Atoi(maxYear)
			g.filterDateCreation(nb, nb2)
		
		}
		minFirstAlbumYear := r.Form.Get("minFirstAlbumYear")
		maxFirstAlbumYear := r.Form.Get("maxFirstAlbumYear")
		if len(minFirstAlbumYear) > 0 && len(maxFirstAlbumYear) > 0 {
			nbr, _ := strconv.Atoi(minFirstAlbumYear)
			nbr2, _ := strconv.Atoi(maxFirstAlbumYear)
			g.filterFirstAlbum(nbr, nbr2)
		}
		maxNumberMembers := r.Form.Get("maxNumberMembers")
		minNumberMembers := r.Form.Get("minNumberMembers")
		if len(minNumberMembers) > 0 && len(maxNumberMembers) > 0 {
			nbmr, _ := strconv.Atoi(minNumberMembers)
			nbmr2, _ := strconv.Atoi(maxNumberMembers)
			g.numberMembers(nbmr, nbmr2)
		}
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


