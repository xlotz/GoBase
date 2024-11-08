package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**
互斥锁
*/

var wait4 sync.WaitGroup
var count3 = 0
var lock3 sync.Mutex

func main() {
	wait4.Add(10)
	for i := 0; i < 10; i++ {
		go func(data *int) {
			// 加锁
			lock3.Lock()
			// 模拟访问耗时
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			// 访问数据
			temp := *data
			// 模拟计算耗时
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			ans := 1
			// 修改数据
			*data = temp + ans

			// 解锁
			lock3.Unlock()
			fmt.Println(*data)
			wait4.Done()
		}(&count3)
	}
	wait4.Wait()
	fmt.Println("最终结果", count3)
}
