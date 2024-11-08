package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("格式化时间...")
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05 Monday Jan"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05 PM"))

	fmt.Println("解析时间...")

	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}

	inLocation, err := time.ParseInLocation("2006/01/02", "2012/10/12", location)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(inLocation.String())

	fmt.Println("Timer...")
	timer := time.NewTimer(time.Second)
	defer timer.Stop()
	select {
	case t := <-timer.C:
		fmt.Println(t)
	}

	fmt.Println("Ticker...")
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for i := 0; i < 3; i++ {
		select {
		case t := <-ticker.C:
			fmt.Println(t)
		}
	}

	fmt.Println("sleep ...")
	start := time.Now()
	fmt.Println(start)
	time.Sleep(time.Second * 2)
	end := time.Now()
	fmt.Println(end)

}
