package main

import "fmt"

type foo int

type person struct {
	first string
	last  string
	age   int
}

func main() {
	var myAge foo
	myAge = 44
	fmt.Printf("%T %v\n", myAge, myAge)

	var yourAge int
	yourAge = 29
	fmt.Printf("%T %v\n", yourAge, yourAge)

	fmt.Println(int(myAge) + yourAge)

	p1 := person{"a", "b", 20}
	p2 := person{"c", "d", 19}
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}
