package main

import (
	"fmt"
	"slices"
)

func main() {

	var num1 [5]int
	num2 := [5]int{1, 2, 3}
	num3 := new([5]int)
	fmt.Println(num1, num2, num3)
	fmt.Println()

	fmt.Println(num1[0])
	num1[0] = 1
	fmt.Println(len(num1), cap(num1))

	nums := [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums[1:])
	fmt.Println(nums[:5])
	fmt.Println(nums[2:3])
	fmt.Println(nums[1:3])

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", arr)
	fmt.Printf("%T\n", arr[1:2])

	slice := slices.Clone(arr[:])
	slice[0] = 0
	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("slice: %v\n", slice)
}
