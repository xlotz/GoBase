package main

import (
	"fmt"
)

func main() {
	h1, e1 := half1(5)
	fmt.Println(h1, e1)

	h2, e2 := half2(5)
	fmt.Println(h2, e2)

	h3 := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}
	fmt.Println(h3(5))

	gre := max1(4, 7, 9, 123, 543, 23, 435, 53, 125)
	fmt.Println(gre)

	// bool 运算
	fmt.Println((true && false) || (false && true) || !(false && false))

	//
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()

}

func half1(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func half2(n int) (float64, bool) {
	return float64(n) / 2, n%2 == 0
}

func max1(numbers ...int) int {
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func foo(numbers ...int) {
	fmt.Println(numbers)
}
