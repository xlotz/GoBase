package main

import (
	e "0_Hello/example"
	"fmt"
	"slices"
)

//var (
//		Stdin = os.NewFile(uintptr(syscall.Stdin), "./stdin.out")
//		Stdout = os.NewFile(uintptr(syscall.Stdout), "./stdout.out")
//		Stderr = os.NewFile(uintptr(syscall.Stderr), "./stderr.out")
//)

func main() {
	//fmt.Println("0_Hello, World!")
	e.SayHello()

	//fmt.Println(math.MaxInt64)
	//os.Stdout.WriteString("0_Hello writer")

	//var buf [1024]byte
	//n, _ := os.Stdin.Read(buf[:])
	//os.Stdout.Write(buf[:n])
	//var a, b int
	//fmt.Scanln(&a, &b)
	//fmt.Printf("%d + %d = %d\n", a, b, a+b)

	/**
	数组和切片
	*/

	// 数组：定长，值类型，值拷贝
	var a [5]int
	a1 := [5]int{1, 2, 3} // 值
	a2 := new([5]int)     // 指针
	fmt.Println(a[0])
	a[0] = 6
	fmt.Println(a)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", arr)
	fmt.Printf("%T\n", arr[1:2])

	// 将数组转换成切片，修改切片会导致修改原数组
	slice := arr[:]
	slice[0] = 0
	fmt.Println(slice)
	fmt.Println(arr)

	// 对切片进行修改,但不改变原数组
	slice2 := slices.Clone(arr[:])
	slice2[0] = 6
	fmt.Println(arr)
	fmt.Println(slice2)

	// 切片： 不定长，会频繁修改和删除
	var nums []int //值
	//nums1 := []int{1,2,3} //值
	//nums2 := make([]int, 0,0) // 值
	//nums3 := new([]int) // 指针

	nums = []int{1, 2, 3, 4, 5}
	nums = append([]int{-1, 0}, nums...)
	fmt.Println(nums)
	i := 3
	nums = append(nums[:i+1], append([]int{999, 999}, nums[i+1:]...)...)
	fmt.Println(nums)
	nums = append(nums, 99, 100)
	fmt.Println(nums)

	// del
	//n := 4
	//nums = nums[n:]
	//fmt.Println(nums)
	//nums = append(nums[:i], nums[i+n:]...)
	//fmt.Println(nums)
	//
	//nums = nums[:0]
	//fmt.Println(nums)

	// copy, 需要足够的长度

	dest := make([]int, 0)
	src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(src, dest)
	fmt.Println(copy(dest, src))
	fmt.Println(src, dest)

	dest1 := make([]int, 10)

	fmt.Println(src, dest1)
	fmt.Println(copy(dest1, src))
	fmt.Println(src, dest1)

	// 遍历
	for i := 0; i < len(src); i++ {
		fmt.Println(src[i])
	}

	for i, v := range src {
		fmt.Println(i, v)
	}

	// 多维切片
	var nu [5][5]int
	for i, num := range nu {
		fmt.Println(i, num)
	}
	fmt.Println()
	sl := make([][]int, 5)
	for i, v := range sl {
		fmt.Println(i, v)
	}
	// 切片的长度不固定，需要单独初始化
	for i := 0; i < len(sl); i++ {
		sl[i] = make([]int, 5)
	}
	for i, v := range sl {
		fmt.Println(i, v)
	}
}

/*
	标识符

	break	default	func	interface	select
	case	defer	go	map	struct
	chan	else	got	package	switch
	const	fallthrough	if	range type
	continue	for	import	return	var

	零值
	数字	0
	布尔	false
	字符串	""
	数组	固定长度的对应类型的零值集合
	结构体	内部字段都是0值的结构体
	切片，映射表，函数，接口，通道，指针 nil

	常量
	const name string = "jack"
	const iota = 0

	变量
	var intNum int
	var str string
	var char byte
	赋值
	第一种：
	var name string = "jack"
	第二种
	var name string
	var age int
	name, age = "jack", 21
	第三种
	name := "jack"

*/
