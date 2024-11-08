package main

import (
	"fmt"
	"sync"
)

/**
Once
*/

var wait7 sync.WaitGroup

func main() {
	var slice MySlice
	wait7.Add(4)
	for i := 0; i < 4; i++ {
		go func() {
			slice.Add(1)
			wait7.Done()
		}()
	}
	wait7.Wait()
	fmt.Println(slice.Len())
}

type MySlice struct {
	s []int
	o sync.Once
}

func (m *MySlice) Get(i int) (int, bool) {
	if m.s == nil {
		return 0, false
	} else {
		return m.s[i], true
	}

}

func (m *MySlice) Add(i int) {
	// 当真正用到切片时，才会考虑初始化

	m.o.Do(func() {
		fmt.Println("初始化...")
		if m.s == nil {
			m.s = make([]int, 0, 10)
		}
	})
	m.s = append(m.s, i)

}

func (m *MySlice) Len() int {
	return len(m.s)

}
