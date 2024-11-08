package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

type doubleZero struct {
	person
	First   string
	License bool
}

func main() {
	p1 := doubleZero{
		person: person{
			First: "A",
			Last:  "B",
			Age:   12,
		},
		First:   "Double zero seven",
		License: true,
	}
	p2 := doubleZero{
		person: person{
			First: "AA",
			Last:  "BB",
			Age:   22,
		},
		First:   "if zero seven",
		License: false,
	}

	fmt.Println(p1.First, p1.person.First)
	fmt.Println(p2.First, p2.person.First)

}
