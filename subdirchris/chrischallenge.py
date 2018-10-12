import requests
import json
import swapi

ships = swapi.get_all("starships")

for ship in ships.iter():
	if ship.pilots:
		print(ship.name)
		for pilots in ship.pilots:
			jason = requests.get(pilots)
			data = jason.json()
			print('   '+data['name'].encode('utf-8'))
