package main

import (
	"testing"
)

func TestGetNumberCities(t *testing.T) {
	got := getNumberCities()
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
	return
}

func TestGetNumberAliens(t *testing.T) {
	got := getNumberAliens()
	want := 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
	return
}

func TestGetRandomCity(t *testing.T) {
	got := getRandomCity()
	want := 0

	if got < want {
		t.Errorf("got %q, wanted %q", got, want)
	}
	return
}

func TestCityExist(t *testing.T) {
	got := cityExist(120)
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
	return
}
