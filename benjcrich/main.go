package main

import (
	"fmt"

	"github.com/peterhellberg/swapi"
)

func main() {
	c := swapi.DefaultClient

	if ship, err := c.Starship(10); err == nil {
		fmt.Println("Name: ", ship.Name)
		p := ship.PilotURLs
		pilots := getPilotName(p)
	}

	// for i := 1; i < 20; i++ {
	// 	if ship, err := c.Starship(10); err == nil {
	// 		fmt.Println("Name: ", ship.Name)
	// 		fmt.Println("Pilots: ", ship.PilotURLs)
	// 	}
	// }
}

func getPilotName(p string) (string) {
	
}
