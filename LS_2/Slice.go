package main

import (
	"fmt"
)

func main() {
	mySlice := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(mySlice)

	fmt.Println(mySlice[1:4])

	// missing low index implies 0
	fmt.Println(mySlice[:3])

	// missing high index implies len(s)
	fmt.Println(mySlice[4:])
}
