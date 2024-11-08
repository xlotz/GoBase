package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var ints byte = 1
	fmt.Println(unsafe.Sizeof(ints))

	var floats float32 = 1.0
	fmt.Println(unsafe.Sizeof(floats))

	var complexs complex128 = 1 + 2i
	fmt.Println(unsafe.Sizeof(complexs))

	var slice []int = make([]int, 100)
	fmt.Println(unsafe.Sizeof(slice))

	var mp map[string]int = make(map[string]int, 0)
	fmt.Println(unsafe.Sizeof(mp))

	type person struct {
		name string
		age  int
	}

	fmt.Println(unsafe.Sizeof(person{}))
	type man struct {
		name string
	}
	fmt.Println(unsafe.Sizeof(man{}))

	fmt.Println("Offsetof...")
	p := person{
		name: "Aa",
		age:  11,
	}
	fmt.Println(unsafe.Sizeof(p))
	fmt.Println(unsafe.Offsetof(p.name))
	fmt.Println(unsafe.Sizeof(p.name))
	fmt.Println(unsafe.Offsetof(p.age))
	fmt.Println(unsafe.Sizeof(p.age))

	fmt.Println("Alignof ...")
	fmt.Println(unsafe.Alignof(p), unsafe.Sizeof(p))
	fmt.Println(unsafe.Alignof(p.name), unsafe.Sizeof(p.name))
	fmt.Println(unsafe.Alignof(p.age), unsafe.Sizeof(p.age))

	fmt.Println("Pointer...")
	fmt.Println("SliceData...")
	nums := []int{1, 2, 3, 4}
	for p, i := unsafe.Pointer(&nums[0]), 0; i < len(nums); p, i = unsafe.Add(p, unsafe.Sizeof(nums[0])), i+1 {
		//for p, i:= unsafe.Pointer(&nums[0]), 0; i<len(nums); p, i= unsafe.Add(p, unsafe.Sizeof(nums[0])), i+1 {
		num := *(*int)(p)
		fmt.Println(num)
	}

	//for p, i := unsafe.Pointer(unsafe.SliceData(nums)), 0; i < len(nums); p, i = unsafe.Add(p, unsafe.Sizeof(int(0))), i+1 {
	//	num := *(int)(p)
	//	fmt.Println(num)
	//}

	numsRef1 := unsafe.Slice(unsafe.SliceData(nums), len(nums))
	numsRef1[0] = 2
	fmt.Println(nums)

	str := "Hello World"
	for ptr, i := unsafe.Pointer(unsafe.StringData(str)), 0; i < len(str); ptr, i = unsafe.Add(ptr, unsafe.Sizeof(byte(0))), i+1 {
		char := *(*byte)(ptr)
		fmt.Println(string(char))
	}

	bytes := []byte("hello world")
	str1 := unsafe.String(unsafe.SliceData(bytes), len(bytes))
	fmt.Println(str1)

}
