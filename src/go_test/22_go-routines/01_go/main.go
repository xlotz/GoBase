package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	fmt.Println("no go ------")
	//foo1()
	//bar1()

	fmt.Println("go concurrency ---------")
	//go foo1()
	//go bar1()

	fmt.Println("wait group ---------")
	//wg.Add(2)
	//go foo2()
	//go bar2()
	//wg.Wait()

	fmt.Println("time sleep ---------")
	//wg.Add(2)
	//go foo3()
	//go bar3()
	//wg.Wait()

	fmt.Println("go maxprocs ---------")
	wg.Add(2)
	go foo3()
	go bar3()
	wg.Wait()
}
func foo1() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo: ", i)
	}
}

func bar1() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar: ", i)
	}
}

func foo2() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo: ", i)
	}
	wg.Done()
}

func bar2() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar: ", i)
	}
	wg.Done()
}

func foo3() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo: ", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}

func bar3() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar: ", i)
		time.Sleep(20 * time.Millisecond)
	}
	wg.Done()
}
