package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 协程
	go fmt.Println("协程1")
	for i := 0; i < 10; i++ {
		go hello(i)
		time.Sleep(time.Millisecond)
	}

	go func() {
		fmt.Println("协程匿名函数")
	}()

	fmt.Println("start")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond)
	}
	time.Sleep(time.Millisecond)
	fmt.Println("end")

}
func hello(i int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	fmt.Println("协程：函数,", i)
}
