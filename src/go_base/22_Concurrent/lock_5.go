package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**
条件变量
*/

var wait6 sync.WaitGroup
var count6 = 0
var rw6 sync.RWMutex

var cond = sync.NewCond(rw6.RLocker())

func main() {

	wait6.Add(12)
	// 读多写少
	go func() {
		for i := 0; i < 3; i++ {
			go Write6(&count6)
		}
		wait6.Done()
	}()

	go func() {
		for i := 0; i < 7; i++ {
			go Read6(&count6)
		}
		wait6.Done()
	}()
	// 等待子协程结束
	wait6.Wait()
	fmt.Println("最终结果", count6)
}

func Read6(i *int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	rw6.RLock()
	fmt.Println("拿到读锁...")
	// 条件不满足就一直阻塞
	for *i < 3 {
		cond.Wait()
	}
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	fmt.Println("释放读锁...", *i)
	rw6.RUnlock()
	wait6.Done()
}

func Write6(i *int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	rw6.Lock()
	fmt.Println("拿到写锁...")
	temp := *i
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	*i = temp + 1
	fmt.Println("释放写锁...", *i)
	rw6.Unlock()
	// 唤醒所有因条件变量阻塞的协程
	cond.Broadcast()
	wait6.Done()
}
