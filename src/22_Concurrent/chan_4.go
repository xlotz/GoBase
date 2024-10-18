package main

import "fmt"

/**
注意点案例
*/

func main() {

	// 以下会导致管道阻塞的情况：

	fmt.Println("1. 读写无缓冲管道")
	//ch1 := make(chan int)
	//defer close(ch1)
	//ch1 <- 123
	//ints, ok := <-ch1
	//fmt.Println(ints, ok)

	fmt.Println("2. 读取空缓冲区的管道")
	//ch2 := make(chan int, 1)
	//defer close(ch2)
	//ints2, ok := <-ch2
	//fmt.Println(ints2, ok)

	fmt.Println("3. 写入满缓冲区的管道")
	//ch3 := make(chan int, 1)
	//defer close(ch3)
	//ch3 <- 1
	//ch3 <- 2

	fmt.Println("4. 管道为nil")
	//var intW chan int
	//intW <- 1
	//
	//var intR chan int
	//fmt.Println(<-intR)
	fmt.Println("4. 关闭一个nil管道")
	//var intC chan int
	//close(intC)

	fmt.Println("5. 写入已关闭的管道")
	//ch4 := make(chan int, 1)
	//close(ch4)
	//ch4 <- 1

	fmt.Println("6. 关闭已关闭的管道")
	ch5 := make(chan int, 1)
	defer close(ch5)
	go write(ch5)
	fmt.Println(<-ch5)
}

func write(ch chan<- int) {
	ch <- 1
	close(ch)
}
