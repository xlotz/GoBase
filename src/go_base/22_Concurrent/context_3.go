package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var waitGroup3 sync.WaitGroup

func main() {
	waitGroup3.Add(3)
	ctx, cancelFunc := context.WithCancel(context.Background())
	go HttpHandler(ctx)
	time.Sleep(time.Second)
	cancelFunc()
	waitGroup3.Wait()
}

func HttpHandler(ctx context.Context) {
	cancelCtxAuth, cancelAuth := context.WithCancel(ctx)
	cancelCtxMail, cancelMail := context.WithCancel(ctx)

	defer cancelAuth()
	defer cancelMail()
	defer waitGroup3.Done()

	go AuthService(cancelCtxAuth)
	go MailService(cancelCtxMail)

	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println("正在处理http请求...")
		}
		time.Sleep(time.Millisecond * 200)
	}
}

func AuthService(ctx context.Context) {
	defer waitGroup3.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("auth 父级取消", ctx.Err())
			return
		default:
			fmt.Println("auth ...")
		}
		time.Sleep(time.Millisecond * 200)
	}
}

func MailService(ctx context.Context) {
	defer waitGroup3.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("mail 父级取消", ctx.Err())
			return
		default:
			fmt.Println("mail ...")
		}
		time.Sleep(time.Millisecond * 200)
	}

}
