package models

type Hero struct {
	ID              int
	Name            string
	Friends         []int
	Characteristics string
	Status          string
}

type SuperHero struct {
	Id     int
	Name   string
	Status string
}
