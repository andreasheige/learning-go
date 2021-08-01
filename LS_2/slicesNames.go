package main

import "fmt"

func main() {
	names := [4]string{
		"Johnny",
		"Conny",
		"Sonny",
		"Ponny",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX" // value at zeroth index of slice b changed
	fmt.Println(a, b)
	fmt.Println(names)
}
