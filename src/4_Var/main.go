package main

import "fmt"

func main() {
	var Num int
	var str string
	var char byte
	fmt.Println(Num, str, char)

	var (
		name string
		age  int
		addr string
	)

	fmt.Println(name, age, addr)

	var N1 string
	N1 = "jack"

	var N2 string = "jack"
	var N3, Age = "jack", 18
	fmt.Println(N1, N2, N3, Age)

	M1 := "tom"
	fmt.Println(M1)

	num1, num2, num3 := 1, 2, 3
	num1, num2, num3 = num3, num1, num2

}
