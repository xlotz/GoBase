package main

import "fmt"

type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person
	License bool
}

func (p person) Greeting() {
	fmt.Println("I'am person.")
}
func (dz doubleZero) Greeting() {
	fmt.Println("Miss Moneypenny")
}

func main() {
	p1 := person{
		Name: "I'm flee",
		Age:  14,
	}
	p2 := doubleZero{
		person: person{
			Name: "Jom Bond",
			Age:  23,
		},
		License: true,
	}
	p1.Greeting()
	p2.Greeting()
	p2.person.Greeting()
}
