package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wait sync.WaitGroup

func main() {
	deadline, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer cancel()
	wait.Add(1)
	go func(ctx context.Context) {
		defer wait.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("上下文取消", ctx.Err())
				return
			default:
				fmt.Println("等待取消中...")
			}
			time.Sleep(time.Millisecond * 200)
		}
	}(deadline)
	wait.Wait()
}
