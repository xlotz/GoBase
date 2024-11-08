package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (z square) area() float64 {
	return z.side * z.side
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(z shape) {
	fmt.Println(z, z.area())
}

// a new method which takes the INTERFACE TYPE shape
func totalArea(shapes ...shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	s := square{10}
	c := circle{5}
	fmt.Printf("%T \n", s)
	info(s)
	fmt.Printf("%T \n", c)
	info(c)

	fmt.Println("Total Area: ", totalArea(c, s))
}
