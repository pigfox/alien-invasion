package main

type Invaders struct {
	alien map[int]Alien
}

type Alien struct {
	id            int
	moves         int
	currentCityId int
}

type Battlefield struct {
	cities map[int]City
}

type City struct {
	id       int
	name     string
	north    string
	west     string
	south    string
	east     string
	occupier int
}
