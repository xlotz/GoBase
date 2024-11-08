package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wait8 sync.WaitGroup

// 临时对象池
var pool sync.Pool

// 用于计数过程中总共创建了多少个对象
var numOfObject atomic.Int64

// BigMemData 假设这是一个占用内存很大的结构体
type BigMemData struct {
	M string
}

func main() {
	pool.New = func() any {
		numOfObject.Add(1)
		return BigMemData{"大内存"}
	}
	wait8.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			val := pool.Get()
			_ = val.(BigMemData)
			pool.Put(val)
			wait8.Done()
		}()
	}
	wait8.Wait()
	fmt.Println(numOfObject.Load())
}
