package main

import "fmt"

type Stringer interface {
	String() string
}

type fakeString struct {
	content string
}

// func used to implement the Stringer interface
func (s *fakeString) String() string {
	return s.content
}

func printString(value interface{}) {
	switch str := value.(type) {
	case string:
		fmt.Println(str)
	case Stringer:
		fmt.Println(str.String())
	}
}

func main() {
	s := &fakeString{"Sailor man come take my hand"}
	printString(s)
	printString("Hello Go!")
}