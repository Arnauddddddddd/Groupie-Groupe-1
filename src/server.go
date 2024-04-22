package engine

import (
	"html/template"
	"net/http"
)

func (g *Structure) Server() {
	http.HandleFunc("/", g.index)
	http.HandleFunc("/index", g.index)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func (g *Structure) index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("pages/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		return
	}
}
