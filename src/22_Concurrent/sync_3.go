package main

import (
	"fmt"
	"sync"
)

func main() {
	var syncMap sync.Map
	// 存入数据
	syncMap.Store("a", 1)
	syncMap.Store("a", "a")
	//读取数据
	fmt.Println(syncMap.Load("a"))
	//读取并删除
	fmt.Println(syncMap.LoadAndDelete("a"))
	// 读取并存入
	fmt.Println(syncMap.LoadOrStore("a", "hello world"))
	syncMap.Store("b", "goodbye world")
	// 变量map
	syncMap.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})

}
