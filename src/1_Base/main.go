package main

import (
	"fmt"
	"math"
)

func SayHello() {
	fmt.Println("SayHello")
}

func sayHello() {
	fmt.Println("sayHello")
}

func main() {
	fmt.Println("Hello")
	fmt.Println(math.MaxInt64)
	SayHello()
	sayHello()
}
