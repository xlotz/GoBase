package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wait2 sync.WaitGroup
var count = 0

func main() {

	wait2.Add(10)
	for i := 0; i < 10; i++ {
		go func(data *int) {
			// 模拟访问耗时
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(5000)))
			// 访问数据
			temp := *data

			// 模拟计算耗时
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(5000)))
			ans := 1
			// 修改数据
			*data = temp + ans
			fmt.Println(*data)
			wait2.Done()
		}(&count)
	}
	wait2.Wait()
	fmt.Println("最终结果", count)
}
