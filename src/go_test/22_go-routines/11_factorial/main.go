package main

import "fmt"

func main() {
	f := factorial(4)
	fmt.Println(f)

	f2 := factorial2(4)
	for n := range f2 {
		fmt.Println(n)
	}
}

func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func factorial2(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}
