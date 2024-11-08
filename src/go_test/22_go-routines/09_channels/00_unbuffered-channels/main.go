package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	// 写
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	// 读
	go func() {
		for {
			fmt.Println(<-c)
		}
	}()
	time.Sleep(time.Second)
}
