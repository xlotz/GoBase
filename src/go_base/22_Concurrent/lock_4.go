package main

/**
读写互斥锁
*/

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wait5 sync.WaitGroup
var count4 = 0
var rw sync.RWMutex

func main() {
	wait5.Add(12)
	// 读多写少
	go func() {
		for i := 0; i < 3; i++ {
			go Write(&count4)
		}
		wait5.Done()
	}()
	go func() {
		for i := 0; i < 7; i++ {
			go Read(&count4)
		}
		wait5.Done()
	}()
	// 等待子协程结束
	wait5.Wait()
	fmt.Println("最终结果", count4)
}

func Read(i *int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	rw.RLock()
	fmt.Println("拿到读锁...")
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	fmt.Println("释放读锁...", *i)
	rw.RUnlock()
	wait5.Done()
}

func Write(i *int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	rw.Lock()
	fmt.Println("拿到写锁...")
	temp := *i
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	*i = temp + 1
	fmt.Println("释放写锁...", *i)
	rw.Unlock()
	wait5.Done()
}
