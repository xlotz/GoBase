package main

import "fmt"

// 死锁
//func main() {
//	c := make(chan int)
//	c <- 1
//	fmt.Println(<-c)
//}

// 解决方案, 修改版
func main() {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)
}
