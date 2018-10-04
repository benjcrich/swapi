#!/usr/bin/env python
import swapi
import requests

ships = (swapi.get_all("starships"))
print (ships)

for ship in ships.iter():
	if len(ship.pilots) > 0:
		print(ship.name)
		for pilot in ship.pilots:
			response = requests.get(pilot)
			data = response.json()
			print("	"+data["name"].encode('utf-8'))