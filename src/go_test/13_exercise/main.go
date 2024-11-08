package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")

	name1 := "Todd"
	fmt.Println("hello, ", name1)

	var name2 string
	fmt.Print("Please enter your name: ")
	fmt.Scan(&name2)
	fmt.Println("Hello, ", name2)

	var numOne, numTwo int
	fmt.Print("Please enter a large number: ")
	fmt.Scan(&numOne)
	fmt.Print("Please enter a smaller number: ")
	fmt.Scan(&numTwo)
	fmt.Println(numOne, "%", numTwo, " = ", numOne%numTwo)

	fmt.Println("even numbers ...")
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("fizzBuzz ...")
	for i := 0; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println(i, " -- FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, " -- FIZZ")
		} else if i%5 == 0 {
			fmt.Println(i, " -- BUZZ")
		} else {
			fmt.Println(i)
		}
	}

	fmt.Println("threeFive ...")
	counter := 0
	for i := 0; i < 100; i++ {
		if i%3 == 0 {
			counter += i
		} else if i%5 == 0 {
			counter += i
		}
	}
	fmt.Println(counter)
}
