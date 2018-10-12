import requests
import json

def getStarships(url):
    r = requests.get(url)
    data = json.loads(r.text)
    return data

def getPilot(piloturl):
    r = requests.get(piloturl)
    pilot = json.loads(r.text)
    name = pilot['name']
    return name

def extractAndPrint(data):
    ships = data['results']
    for x in range(len(ships)):
        ship = ships[x]
        if bool(ship['pilots']) != False:
            print("The " + ship['name'] + " was piloted by...")
            pilots = ship['pilots']
            for i in range(len(pilots)):
                piloturl =  pilots[i]
                print("\t" + getPilot(piloturl))


url = 'https://swapi.co/api/starships/'

while url != "":
    try:
        data = getStarships(url)
        url = data['next']
        extractAndPrint(data)
    except:
        quit()
