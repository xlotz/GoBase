package main

import "fmt"

const (
	a1 = iota
	b1 = iota
	c1 = iota
)

const (
	a2 = iota
	b2
	c2
)

const (
	_  = iota
	b3 = iota * 10
	c3 = iota * 10
)

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
	TB = 1 << (iota * 10)
)

func main() {
	fmt.Println(a1, b1, c1)
	fmt.Println(a2, b2, c2)
	fmt.Println(b3, c3)
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t, %d\n", KB, KB)
	fmt.Printf("%b\t, %d\n", KB, KB)
	fmt.Printf("%b\t, %d\n", MB, MB)
	fmt.Printf("%b\t, %d\n", MB, MB)
	fmt.Printf("%b\t, %d\n", TB, TB)
	fmt.Printf("%b\t, %d\n", TB, TB)
}
