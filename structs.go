package main

type Invaders struct {
	alien map[int]Alien
}

type Alien struct {
	moves         int
	currentCityId int
}

type Battlefield struct {
	cities map[int]City
}

type City struct {
	name     string
	north    string
	west     string
	south    string
	east     string
	occupier int
}
