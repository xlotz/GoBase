package main

import (
	"fmt"
	"slices"
)

type Person struct {
	Name   string
	Age    int
	Salary float64
}

func main() {
	var sum2 = func(a, b int) int {
		return a + b
	}
	fmt.Println(sum1(1, 2))
	fmt.Println(sum2(1, 2))
	fmt.Println(sum3(1, 2))
	fmt.Println(sum4(1, 2))

	// 匿名函数
	a := func(a, b int) int {
		return a + b
	}(2, 3)
	fmt.Println(a)

	//
	people := []Person{
		{Name: "a", Age: 25, Salary: 5000.0},
		{Name: "b", Age: 30, Salary: 6000.0},
		{Name: "c", Age: 28, Salary: 5500.0},
	}
	slices.SortFunc(people, func(p1, p2 Person) int {
		if p1.Name > p2.Name {
			return 1
		} else if p1.Name < p2.Name {
			return -1
		}
		return 0
	})

	//
	grow := Exp(2)
	for i := 0; i < 10; i++ {
		fmt.Printf("2^%d=%d\n", i, grow())
	}

	// 10个斐波那契数
	fib := Fib(10)
	for n, next := fib(); next; n, next = fib() {
		fmt.Println(n)
	}

	// Do
	Do()
	fmt.Println(0)
	Do2()

	fmt.Println(sum5(3, 5))

	// 预计数
	defer fmt.Println(Fn1())
	fmt.Println("3")

	aa, bb := 1, 2
	fmt.Println("aa + bb")
	defer fmt.Println(sum1(aa, bb))
	aa, bb = 3, 4

	f := func() {
		fmt.Println(sum1(aa, bb))
	}
	aa, bb = 3, 4
	f()

	defer func() {
		fmt.Println(sum1(aa, bb))
	}()
	aa, bb = 3, 4
	defer func(num int) {
		fmt.Println(num)
	}(sum1(aa, bb))
	aa, bb = 3, 4
}

func sum1(a, b int) int {
	return a + b
}

func sum3(a, b int) (ans int) {
	ans = a + b
	return
}

func sum4(a, b int) (c, d int) {
	c = a + b
	d = a * b
	return a - b, a + b
}

func Exp(n int) func() int {
	e := 1
	return func() int {
		temp := e
		e *= n
		return temp
	}
}

func Fib(n int) func() (int, bool) {
	a, b, c := 1, 1, 2
	i := 0
	return func() (int, bool) {
		if i >= n {
			return 0, false
		} else if i < 2 {
			f := i
			i++
			return f, true
		}
		a, b = b, c
		c = a + b
		i++
		return a, true
	}
}

// 延迟调用
func Do() {
	defer func() {
		fmt.Println("1")
	}()
	fmt.Println("2")
}

func sum5(a, b int) (s int) {
	defer func() {
		s -= 10
	}()
	s = a + b
	return
}

func Do2() {
	defer fmt.Println(1)
	fmt.Println(2)
	defer fmt.Println(3)
	defer fmt.Println(4)
	fmt.Println(5)
}

// 参数预计算
func Fn1() int {
	fmt.Println("2")
	return 1
}
