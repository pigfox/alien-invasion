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

const alienMoveLimit = 10 //10000

func init() {
	directions = []string{"north", "west", "south", "east"}
	invaders = Invaders{make(map[int]Alien)}
	battlefield = Battlefield{make(map[int]City)}
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
	if 0 < len(invaders.alien) {
		for k, v := range invaders.alien {
			tmp := invaders.alien
			moves := v.moves
			moves++

			tmp[0] = Alien{moves: moves, currentCityId: getRandomOriginCity()}
			invaders.alien[k] = tmp[0]
			//isCityOccupied()
			if alienMoveLimit <= invaders.alien[k].moves {
				removeAlien(k)
			}
		}
	} else {
		gameOver("Game over. All invaders have died...")
	}
}

func setCityOccupied(id int) {
	tmp := battlefield.cities[id]
	tmp.occupier = id
	battlefield.cities[id] = tmp
}

func isCityOccupied() {

}

func useTzarBomba() {

}

func getCityOccupier(id int) int {
	return battlefield.cities[id].occupier
}

func removeAlien(id int) {
	delete(invaders.alien, id)
}

func getRandomOriginCity() int {
	id := -1
	lengthCities := len(battlefield.cities)
	//make sure all cities have not been nuked...
	if 0 < lengthCities {
		id = rand.Intn(lengthCities)
		if cityExist(id) {
			direction := directions[rand.Intn(len(directions))]
			exists := targetCityExists(direction, id)
			if exists {
				return id
			} else {
				getRandomOriginCity()
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
		invaders.alien[m] = Alien{moves: 0, currentCityId: 0}
	}
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
			log.Fatalf("Warning corrupt file line on " + fmt.Sprintf("%d", id))
		}
	}
	//fmt.Println(battlefield)
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
