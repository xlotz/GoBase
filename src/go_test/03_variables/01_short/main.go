package main

import "fmt"

func main() {

	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like go?`
	g := 'M'

	fmt.Printf("%v \t %v \t %v \t %v \t %v \t %v \t %v\n", a, b, c, d, e, f, g)
	fmt.Printf("%T \t %T \t %T \t %T \t %T \t %T \t %T\n", a, b, c, d, e, f, g)

}
