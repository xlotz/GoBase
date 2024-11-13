package main

import "fmt"

//func main() {
//	c := make(chan int)
//	go func() {
//		for i := 0; i < 10; i++ {
//			c <- i
//		}
//	}()
//
//
//	//fmt.Println(<-c)  // 0
//
//	//for {
//	//	fmt.Println(<-c)
//	//}  // fatal error: all goroutines are asleep - deadlock!
//
//}

/**
修改版
*/

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}
}
