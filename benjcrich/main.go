package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	// "reflect"
)

type Response struct {
	Name  string             `json:"name"`
	Next  string             `json:"next"`
	Ships []StarshipResponse `json:"results"`
}

type StarshipResponse struct {
	Name   string   `json:"name"`
	Pilots []string `json:"pilots"`
}

func main() {

	const s = `   .           .        .                     .          .            .
    .               .    .          .              .   .         .
          _________________      ____         __________
    .    /                 |    /    \    .  |          \
.       /     _____   _____| . /      \      |    ___    |     .     .
        \    \    |   |       /   /\   \     |   |___>   |
      .  \    \   |   |      /   /__\   \  . |         _/               .
.    _____>    |  |   | .   /            \   |   |\    \_______    .
    |          /  |   |    /    ______    \  |   | \           |
    |_________/   |___|   /____/      \____\ |___|  \__________|    .
.     ____    __  . _____   ____      .  __________   .  _________
     \    \  /  \  /    /  /    \       |          \    /         |      .
      \    \/    \/    /  /      \      |    ___    |  /    ______|  .
       \              /  /   /\   \ .   |   |___>   |  \    \
.       \            /  /   /__\   \    |         _/.   \    \            +
         \    /\    /  /            \   |   |\    \______>    |   .
          \  /  \  /  /    ______    \  |   | \              /          .
.     .    \/    \/  /____/      \____\ |___|  \____________/  LS
                            .                                        .
  .                           .         .               .                 .
             .                                   .            .
________________________________________________________________________
|:..                                                      :::::%%%%%%HH|
|%%%:::::..        S t a r s h i p s  &  P i l o t s         ::::::%%%%|
|HH%%%%%:::::....._______________________________________________::::::|`

	fmt.Printf("%v\n\n", s)

	url := "https://swapi.dev/api/starships/?page=1"
	starships := get(url)

	for i := 0; i < len(starships.Ships); i++ {
		if len(starships.Ships[i].Pilots) == 0 {
			i++
		} else {
			fmt.Println(starships.Ships[i].Name)

			getPilots(starships, i)
		}
	}
}

func get(url string) Response {
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var starships Response
	json.Unmarshal(responseData, &starships)

	return starships
}

func getPilots(starships Response, i int) {
	p := starships.Ships[i].Pilots
	fmt.Println("Pilots: ")
	for _, url := range p {
		pilot := get(url)
		trimmedPilot := strings.Trim(pilot.Name, "[]{}")
		fmt.Printf("    - %s\n", string(trimmedPilot))
	}
}
