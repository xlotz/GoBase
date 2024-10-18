package main

import "fmt"

/*
*
利用管道的阻塞条件，写一个主协程等待子协程执行完毕的例子
*/
func main() {
	// 创建无缓冲区管道
	ch := make(chan struct{})
	defer close(ch)

	go func() {
		fmt.Println(2)
		// 写入
		ch <- struct{}{}
	}()
	// 阻塞等待读取
	<-ch
	fmt.Println(1)

}
