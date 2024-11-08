package main

import "fmt"

func main() {

	// int
	var x [58]int
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
	x[42] = 777
	fmt.Println(x[42])

	// string
	var x1 [58]string
	for i := 65; i <= 122; i++ {
		x1[i-65] = string(i)
	}
	fmt.Println(x1)
	fmt.Println(x1[42])

	//
	var x2 [256]int
	fmt.Println(len(x2))
	fmt.Println(x2[42])
	for i := 0; i < 256; i++ {
		x2[i] = i
	}
	for i, v := range x2 {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if i > 50 {
			break
		}
	}

	//

	var x3 [256]byte
	fmt.Println(len(x3))
	fmt.Println(x3[42])
	for i := 0; i < 256; i++ {
		x3[i] = byte(i)
	}
	for i, v := range x3 {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if i > 50 {
			break
		}
	}

	//
	var x4 [256]string
	fmt.Println(len(x4))
	fmt.Println(x4[0])
	for i := 0; i < 256; i++ {
		x4[i] = string(i)
	}
	for _, v := range x4 {
		fmt.Printf("%v - %T - %v\n", v, v, []byte(v))
	}
}
