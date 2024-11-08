package main

import (
	"fmt"
	"time"
)

/*
*
演示有缓冲区读写发生阻塞线下
*/
func main() {

	ch := make(chan int, 5)
	// 创建2个无缓冲管道
	chW := make(chan struct{})
	chR := make(chan struct{})

	defer func() {
		close(ch)
		close(chW)
		close(chR)
	}()

	// 负责写
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("写入: ", i)
		}
		chW <- struct{}{}
	}()
	// 负责读
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond)
			fmt.Println("读取, ", <-ch)
		}
		chR <- struct{}{}
	}()

	fmt.Println("写入完成,", <-chW)
	fmt.Println("读取完成,", <-chR)

}
