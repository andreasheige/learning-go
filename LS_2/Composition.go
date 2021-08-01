package main

import (
	"fmt"
)

type User struct {
	Id       int
	Name     string
	Location string
}

//type Player with one additional attribute
type Player struct {
	Id       int
	Name     string
	Location string
	GameId   int
}

func main() {
	p := Player{}
	p.Id = 13
	p.Name = "Andreas"
	p.Location = "GBG"
	p.GameId = 41307
	fmt.Printf("%+v", p) // the value in a default format when printing structs,
	// the plus flag (%+v) adds field names
}
