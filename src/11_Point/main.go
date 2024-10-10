package main

import "fmt"

func main() {

	num := 2
	// 取地址
	p := &num
	fmt.Println(p)

	// 取值
	r := *p
	fmt.Println(r)

	var n *int
	fmt.Println(n)
	// 初始化
	n = new(int)
	fmt.Println(n)

	// 零值
	fmt.Println(*new(string))
	fmt.Println(*new(int))
	fmt.Println(*new([5]int))
	fmt.Println(*new([]float64))

	/**
	new make
	*/
	//new(int) // int指针
	//new(string) // string指针
	//new([]int) // 整型切片指针
	//make([]int, 10, 100) // 长度为10，容量100的整型切片
	//make(map[string]int, 10) // 容量为10的映射表
	//make(chan int, 10) // 缓冲区大小为10的通道

	fmt.Println(new(int))
	fmt.Println(make([]int, 10, 100))
	fmt.Println(make(map[string]int, 10))
	fmt.Println(make(chan int, 10))

}
