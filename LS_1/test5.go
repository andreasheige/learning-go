package main

import "fmt"

type Point struct {
	X, Y int
}

var (
	p = Point{108, 22}  // has type point
	q = &Point{12, 208} // has type *Point
	r = Point{X: 1000}  // Y:0 is implicit
	s = Point{}         // X:0 and Y:0
)

func main() {
	fmt.Println(p, q, r, s)
}
