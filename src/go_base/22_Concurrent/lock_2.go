package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wait3 sync.WaitGroup
var count2 = 0
var lock2 sync.Mutex

func main() {
	wait3.Add(10)
	for i := 0; i < 10; i++ {
		go func(data *int) {
			// 加锁
			lock2.Lock()
			//模拟访问耗时
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			// 访问数据
			temp := *data
			// 模拟计算耗时
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			ans := 1
			// 修改数据
			*data = temp + ans
			// 解锁
			lock2.Unlock()
			fmt.Println(*data)
			wait3.Done()
		}(&count2)
	}
	wait3.Wait()
	fmt.Println("最终结果", count2)
}
