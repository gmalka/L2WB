package model

type User struct {
	Name     string
	Events   map[int]Event
}