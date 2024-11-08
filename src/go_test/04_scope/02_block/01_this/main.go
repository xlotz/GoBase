package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	//foo()
	foo2(x)
}

//func foo() {
//	fmt.Println(x)
//}

func foo2(x int) {
	fmt.Println(x)
}
