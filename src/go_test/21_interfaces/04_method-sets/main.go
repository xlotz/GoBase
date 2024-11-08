package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape1 interface {
	area1() float64
	//area2() float64
}
type shape2 interface {
	//area1() float64
	area2() float64
}

func (c circle) area1() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) area2() float64 {
	return math.Pi * c.radius * c.radius
}

func info1(s shape1) {
	fmt.Println("area", s.area1())
}
func info2(s shape2) {
	fmt.Println("area", s.area2())
}
func main() {
	c := circle{5}
	info1(c)
	info2(&c)
}
