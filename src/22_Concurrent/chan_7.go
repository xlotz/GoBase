package main

import (
	"fmt"
	"time"
)

/**
select
*/

func main() {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)

	defer func() {
		close(chA)
		close(chB)
		close(chC)
	}()

	//go func() {
	//	chA <- 1
	//}()

	go Send(chA)
	//go Send(chB)
	//go Send(chC)

	for {
		select {
		case n, ok := <-chA:
			fmt.Println(n, ok)
		case n, ok := <-chB:
			fmt.Println(n, ok)
		case n, ok := <-chC:
			fmt.Println(n, ok)
		default:
			fmt.Println("所有管道都不可用")
		}
		time.Sleep(time.Minute)
	}

}

func Send(ch chan<- int) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond)
		ch <- i
	}
}
