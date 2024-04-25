package engine

import "fmt"

func (g *Structure) Search() {
	fmt.Println(g.artists[51].Name)
	fmt.Println(g.artists[51].Locations.Locations)
}