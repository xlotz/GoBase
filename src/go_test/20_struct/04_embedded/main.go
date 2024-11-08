package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

type doubleZero struct {
	person
	License bool
}

func main() {
	p1 := doubleZero{
		person: person{
			First: "Tom",
			Last:  "Bond",
			Age:   20,
		},
		License: true,
	}
	p2 := doubleZero{
		person: person{
			First: "Jack",
			Last:  "BOO",
			Age:   19,
		},
		License: false,
	}
	fmt.Println(p1.First, p1.Last, p1.Age, p1.License)
	fmt.Println(p2.First, p2.Last, p2.Age, p2.License)

}
