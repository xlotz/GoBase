package main

import (
	"fmt"
	"sync"
)

func main() {
	// 并发map
	//myMap := make(map[int]int, 10)
	var syncMap2 sync.Map
	var wait10 sync.WaitGroup
	wait10.Add(10)
	for i := 0; i < 10; i++ {
		go func(n int) {
			for i := 0; i < 100; i++ {
				syncMap2.Store(n, n)
			}
			wait10.Done()
		}(i)
	}
	wait10.Wait()
	syncMap2.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}
