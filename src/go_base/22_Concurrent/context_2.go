package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var waitGroup2 sync.WaitGroup

func main() {

	bkg := context.Background()
	// 返回一个cancelCtx 和 cancel函数
	cancelCtx, cancel := context.WithCancel(bkg)
	waitGroup2.Add(1)
	go func(ctx context.Context) {
		defer waitGroup2.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				return
			default:
				fmt.Println("等待取消中...")
			}
			time.Sleep(time.Millisecond * 200)

		}
	}(cancelCtx)
	time.Sleep(time.Second)
	cancel()
	waitGroup2.Wait()
}
