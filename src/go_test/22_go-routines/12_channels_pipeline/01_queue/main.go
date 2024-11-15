package main

import "fmt"

func main() {
	c := gen(2, 3)
	out := squeue(c)

	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func squeue(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out

}
