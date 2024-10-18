package main

import (
	"fmt"
)

/**
for range
*/

func main() {
	ch := make(chan int, 10)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		ch <- i
	//	}
	//	close(ch)
	//}()
	//
	//for n := range ch {
	//	fmt.Println(n)
	//}

	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
	for i := 0; i < 6; i++ {
		n, ok := <-ch
		fmt.Println(n, ok)
	}
}
