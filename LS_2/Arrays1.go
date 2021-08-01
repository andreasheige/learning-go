package main

import (
	"fmt"
)

func main() {
	var a [2]string // array of size 2
	a[0] = "Wello"
	a[1] = "Horld"
	fmt.Println(a[0], a[1]) // will print both arrays
	fmt.Println(a)          // will print both arrays as one
}
