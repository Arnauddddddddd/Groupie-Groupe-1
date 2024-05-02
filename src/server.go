package engine

import (
	"fmt"
	"html/template"
	"net/http"
)

type WebStruct struct {
	Artists []ArtistsStruct
}


func (g *Structure) Server() {
	http.HandleFunc("/", g.index)
	http.HandleFunc("/index.html", g.index)
	http.Handle("/videos/", http.StripPrefix("/videos/", http.FileServer(http.Dir("videos"))))
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))

	fmt.Println("\nhttp://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {return}
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




	web := WebStruct{Artists: g.artistsTemp,}
	err := tmpl.Execute(w, web)
	if err != nil {
		return
	}
}


