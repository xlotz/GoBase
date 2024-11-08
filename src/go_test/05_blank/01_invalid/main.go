package main

import "fmt"

func main() {
	a := "a"
	//b := "b"
	fmt.Println("a - ", a)

	// b 没有用到，编译无法通过
	//fmt.Println(b)
}
