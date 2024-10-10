package main

import "fmt"

type IntSlice []int

func (i IntSlice) Get(index int) int {
	return i[index]
}

func (i IntSlice) Set(index int, val int) {
	i[index] = val
}

func (i IntSlice) Len() int {
	return len(i)
}

type MyInt int

// 值接收者
func (i MyInt) Set(v int) {
	i = MyInt(v) // 修改但不造成任何影响
}

// 指针接收者
func (i *MyInt) Set2(v int) {
	*i = MyInt(v)
}
func main() {

	// 方法
	var intSlice IntSlice
	intSlice = []int{1, 2, 3, 4, 5}
	fmt.Println(intSlice.Get(0))
	intSlice.Set(0, 7)
	fmt.Println(intSlice.Get(0))
	fmt.Println(intSlice.Len())

	// 值接收者
	fmt.Println("值接收者")
	myInt := MyInt(1)
	myInt.Set(2)
	fmt.Println(myInt)

	// 指针接收者
	fmt.Println("指针接收者")
	myInt2 := MyInt(1)
	myInt2.Set2(2)
	fmt.Println(myInt2)

}
