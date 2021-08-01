package main

import (
	"fmt"
)

type User struct {
	Id             int
	Name, Location string
}

func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s", u.Name, u.Location)
}

//type Player with one additional attribute
type Player struct {
	User   //user will contain all the required attributes
	GameId int
}

func main() {
	p := Player{}
	p.Id = 13
	p.Name = "Andreas"
	p.Location = "GBG"
	p.GameId = 41307
	fmt.Println(p.Greetings())
	//fmt.Printf("%+v", p) // the value in a default format when printing structs,
	// the plus flag (%+v) adds field names
}

// func main() {
// 	p := Player{
// 		User{Id: 13, Name: "Andreas", Location: "GBG"},
// 		90404,
// 	}
// 	fmt.Printf(
// 		"Id: %d, Name: %s, Location: %s, Game id: %d\n",
// 		p.Id, p.Name, p.Location, p.GameId)
// 	// Directly set a field defined on the Player struct
// 	p.Id = 11
// 	fmt.Printf("%+v", p)
// }
