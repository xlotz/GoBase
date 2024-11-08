package main

import "fmt"

// 单向管道
func main() {
	ch := make(chan int, 1)
	go write2(ch)
	fmt.Println(<-ch)
}

func write2(ch chan<- int) {
	ch <- 1
}
