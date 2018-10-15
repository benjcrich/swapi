#! /bin/python
import json
import requests
import pprint

ship_url = 'https://swapi.co/api/starships/'
people_url = 'https://swapi.co/api/people/'

pp = pprint.PrettyPrinter()

#API call for the ships, looping or pagination
def get_ships(ship_url):
    ships_results = requests.get(ship_url)
    ship_data = ships_results.json()["results"]
    while ships_results.json()["next"]:
        ships_results = requests.get(ships_results.json()["next"])
        ship_data = ship_data + ships_results.json()["results"]
    print("Gathered ship data")
#Sends the results of all the ships in a single list  and the people_url to get_people
    get_people(people_url, ship_data)

#API call for all the people, looping for pagination
def get_people(people_url, ship_data):
    people_results = requests.get(people_url)
    people_data = people_results.json()["results"]
    while people_results.json()["next"]:
        people_results = requests.get(people_results.json()["next"])
        people_data = people_data + people_results.json()["results"]
    print("Gathered people data \n\r --------------------------------")
#Passes the list of ships and the list of people to list_ships
    list_ships(people_data, ship_data)

#checks if the ships has pilots and prints the name of the ship if it does.
def list_ships(people_data, ship_data):
    for ship in ship_data:
        if ship["pilots"]:
            print(ship["name"])
            #Passes the ship's info to list_pilots as well as the people_data from before
            list_pilots(people_data, ship)

#cross references the data from the ships that list the urls of pilots and prints the name for that entry
def list_pilots(people_data, ship):
    for pilot_url in ship["pilots"]:
        for person in people_data:
            if person["url"] == pilot_url:
                print("  "+person["name"])

if __name__ == "__main__":
    get_ships(ship_url)
