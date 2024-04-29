package engine

import (
	"fmt"
	"html/template"
	"net/http"
)

type WebStruct struct {
	Id           []int
	Artists      []string
	Members      [][]string
	DateCreation []int
	FirstAlbum   []string
	Poster       []string
	Dates        [][]string
}

func (g *Structure) Server() {
	http.HandleFunc("/", g.index)
	http.HandleFunc("/index.html", g.index)
	http.Handle("/videos/", http.StripPrefix("/videos/", http.FileServer(http.Dir("videos"))))
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))

	fmt.Println("\nhttp://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func (g *Structure) index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("pages/index.html"))
	web := WebStruct{
		Id:           g.artistsId,
		Artists:      g.artistsName,
		Members:      g.artistsMembers,
		DateCreation: g.dateCreation,
		FirstAlbum:   g.firstAlbum,
		Poster:       g.artistsPosters,
		Dates:        g.dates,
	}

	err := tmpl.Execute(w, web)
	if err != nil {
		return
	}
}
