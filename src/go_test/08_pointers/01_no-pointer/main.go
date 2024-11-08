package main

import "fmt"

func main() {
	a := 43
	fmt.Println(a, " ", &a)
	var b = &a
	fmt.Println(b)
	fmt.Println(*b) // 43

	// 修改
	*b = 42
	fmt.Println(a) // 42

	// no pointer
	x := 5
	fmt.Printf("%p\n", &x)
	fmt.Println(&x)

	zero(x)
	fmt.Println(x)

}

func zero(z int) {
	fmt.Printf("%p\n", &z)
	fmt.Println(&z)
	z = 0
}
