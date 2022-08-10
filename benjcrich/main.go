package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	// "reflect"
)

type Response struct {
	Name  string             `json:"name"`
	Ships []StarshipResponse `json:"results"`
}

type StarshipResponse struct {
	Name   string   `json:"name"`
	Pilots []string `json:"pilots"`
}

func main() {
	url := "https://swapi.dev/api/starships/?page=2"
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

	for _, url := range p {
		pilot := get(url)
		trimPilot := strings.Trim(pilot, "[]{}")
		fmt.Printf("    %s\n", trimPilot)
	}
}
