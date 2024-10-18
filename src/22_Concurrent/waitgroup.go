package main

import (
	"fmt"
	"sync"
)

func main() {
	var mainWait sync.WaitGroup
	var wait sync.WaitGroup
	// 指定协程的数量
	mainWait.Add(10)
	fmt.Println("Start...")

	for i := 0; i < 10; i++ {
		// 循环计数
		wait.Add(1)
		go func() {

			fmt.Println(i)

			// 两个计数器-1
			mainWait.Done()
			wait.Done()
		}()
		// 等待当前循环的协程执行完毕
		wait.Wait()
	}

	// 等待所有子协程
	mainWait.Wait()
	fmt.Println("end")

}
