package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var invaders Invaders
var battlefield Battlefield
var directions []string

const alienMoveLimit = 10000

func init() {
	directions = []string{"north", "west", "south", "east"}
	invaders = Invaders{make(map[int]Alien)}
	battlefield = Battlefield{map[int]City{}}
	readBattlefield()
}

func main() {
	fmt.Println("Alien invasion starting...")
	fmt.Println("Enter number of aliens you wish to create.")

	createAliens(getInput())
	for {
		invadeRandomCity()
	}
}

func getInput() int {
	input := ""
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err)
	}

	numAliens, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	return numAliens
}

func gameOver(s string) {
	fmt.Println(s)
	os.Exit(0)
}

func invadeRandomCity() {
	if 0 < getNumberAliens() {
		if getNumberCities() == 0 {
			gameOver("Game over! No cities left to invade...")
		}

		for alienID, alien := range invaders.alien {
			moves := alien.moves
			moves++

			cityID := getRandomCity()

			tmpAlien := invaders.alien
			tmpAlien[0] = Alien{moves: moves, currentCityId: cityID}
			invaders.alien[alienID] = tmpAlien[0]
			setCityOccupier(alienID, cityID)
			//checkAliens()
			/*
				fmt.Print(cityID)
				fmt.Print(battlefield.cities[cityID])
				fmt.Print("occupier:")
				fmt.Println(alienID)


					fmt.Print("occupier:")
					fmt.Println(occupier)
					fmt.Print("cityID:")
					fmt.Println(cityID)
					fmt.Print("k:")
					fmt.Println(k)
			*/

			if alienMoveLimit <= invaders.alien[alienID].moves {
				deleteAlien(alienID)
			}
			//fmt.Print("alienID:")
			//fmt.Println(alienID)
			//checkAliens()
			//checkCities()
		}
	} else {
		gameOver("Game over! All aliens have died...")
	}
}

func getNumberCities() int {
	return len(battlefield.cities)
}

func getNumberAliens() int {
	return len(invaders.alien)
}

func setCityOccupier(challengerID int, cityID int) {
	//city is already occupied
	occupierID := getCityOccupier(cityID)
	if -1 < occupierID {
		destroy(occupierID, challengerID, cityID)
		return
	}
	tmpCity := battlefield.cities[cityID]
	tmpCity.occupier = challengerID
	battlefield.cities[cityID] = tmpCity
}

func getCityOccupier(id int) int {
	return battlefield.cities[id].occupier
}

func destroy(occupierID int, challengerID int, cityID int) {
	if occupierID < 0 || challengerID < 0 || cityID < 0 {
		s := "occupierID:" + strconv.Itoa(occupierID) + "\n"
		s += "challengerID:" + strconv.Itoa(challengerID) + "\n"
		s += "cityID:" + strconv.Itoa(cityID) + "\n"
		log.Fatalf(s)
	}

	if occupierID == challengerID {
		return
	}

	city := battlefield.cities[cityID].name
	if city == "" {
		invadeRandomCity()
	}
	a1 := strconv.Itoa(occupierID)
	a2 := strconv.Itoa(challengerID)

	s := city + " has been destroyed by alien " + a1 + " and alien " + a2 + "!"
	fmt.Println(s)

	deleteAlien(occupierID)
	deleteAlien(challengerID)
	deleteCity(cityID)
	invadeRandomCity()
}

func deleteAlien(id int) {
	if 0 < len(invaders.alien) {
		fmt.Println("Deleting alien #" + strconv.Itoa(id))
		delete(invaders.alien, id)
		fmt.Println("Number of aliens left: " + strconv.Itoa(len(invaders.alien)))
	}
}

func deleteCity(id int) {
	if 0 < len(battlefield.cities) {
		fmt.Println("Deleting city #" + strconv.Itoa(id))
		delete(battlefield.cities, id)
		fmt.Println("Number of cities left: " + strconv.Itoa(len(battlefield.cities)))
	}
}

func getRandomCity() int {
	id := -1
	lengthCities := len(battlefield.cities)
	//make sure all cities have not been destroyed...
	if 0 < lengthCities {
		id = rand.Intn(lengthCities)
		if cityExist(id) {
			direction := directions[rand.Intn(len(directions))]
			exists := targetCityExists(direction, id)
			if exists {
				return id
			} else {
				getRandomCity()
			}
		} else {
			fmt.Println("City not found")
		}
	}
	return id
}

func targetCityExists(dir string, id int) bool {
	newCity := ""
	switch dir {
	case "north":
		newCity = battlefield.cities[id].north
	case "west":
		newCity = battlefield.cities[id].west
	case "south":
		newCity = battlefield.cities[id].south
	case "east":
		newCity = battlefield.cities[id].east
	}

	for _, v := range battlefield.cities {
		if v.name == newCity {
			return true
		}
	}

	return false
}

func cityExist(id int) bool {
	t := reflect.TypeOf(battlefield.cities[id]).Kind()
	c := reflect.TypeOf(City{}).Kind()
	if reflect.TypeOf(c) != reflect.TypeOf(t) {
		return false
	}
	return true
}

func createAliens(n int) {
	for m := 0; m < n; m++ {
		invaders.alien[m] = Alien{id: m, moves: 0, currentCityId: 0}
	}
	//checkAliens()
}

func readBattlefield() {
	fileName := "./map.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open " + fileName)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()
	for id, cityLine := range text {
		cityChunks := strings.Split(cityLine, " ")
		//Check for valid line
		if len(cityChunks) == 5 {
			city := City{}
			city.id = id
			city.name = cityChunks[0]

			north := strings.Split(cityChunks[1], "=")
			city.north = north[1]

			west := strings.Split(cityChunks[2], "=")
			city.west = west[1]

			south := strings.Split(cityChunks[3], "=")
			city.south = south[1]

			east := strings.Split(cityChunks[4], "=")
			city.east = east[1]

			city.occupier = -1

			battlefield.cities[id] = city
		} else {
			log.Fatalf("Error corrupt file line on " + strconv.Itoa(id))
		}
	}
}

//Function to verify that map indices are consistent with city ids
func checkCities() {
	fmt.Println("Cities start")
	for k, v := range battlefield.cities {
		fmt.Print(k)
		fmt.Print(":")
		fmt.Println(v)
		if k != v.id {
			s := "City Error corrupt ordering, index "
			s += strconv.Itoa(k) + " != id " + strconv.Itoa(v.id)
			log.Fatalf(s)
		}
	}
	fmt.Println("Cities end")
}

//Function to verify that map indices are consistent with alien ids
func checkAliens() {
	fmt.Println("Aliens start")
	for k, v := range invaders.alien {
		fmt.Print(k)
		fmt.Print(":")
		fmt.Println(v)
		if k != v.id {
			s := "Alien Error corrupt ordering, index "
			s += strconv.Itoa(k) + " != id " + strconv.Itoa(v.id)
			log.Fatalf(s)
		}
	}
	fmt.Println("Aliens end")
}
