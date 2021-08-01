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

	primes := [6]int{2, 3, 5, 7, 11, 13} // array of prime numbers of size 6
	fmt.Println(primes)

	b := [2]string{"tjabba", "tjena"}
	fmt.Printf("%q", b)
}
