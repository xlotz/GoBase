package main

import "fmt"

func main() {
	intCh := make(chan int, 1)
	defer close(intCh)

	intCh <- 1
	//
	//fmt.Println(<-intCh)
	// 读取有两个返回值
	ints, ok := <-intCh
	fmt.Println(ints, ok)

	// 死锁
	fmt.Println("演示死锁...")
	ch1 := make(chan int)
	defer close(ch1)
	//ch1 <- 123
	//n := <-ch1
	//fmt.Println(n) // 0缓冲区的管道写入数据时必须有其他协程立即读取，否则会发生死锁
	// 修改为
	go func() {
		ch1 <- 123
	}()

	n2 := <-ch1
	fmt.Println(n2)

	fmt.Println("有缓冲区")
	ch2 := make(chan int, 1)
	defer close(ch2)
	ch2 <- 121
	n3 := <-ch2
	fmt.Println(n3)

}
