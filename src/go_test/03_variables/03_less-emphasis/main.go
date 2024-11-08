package main

import "fmt"

// 全局配置
var a = "this is stored"
var b, c string = "stored b", "stored c"
var d string
var name = "Todd"

func main() {

	// var
	var message string
	message = "hello"
	fmt.Println(message)

	//many
	var a1, b1, c1 int
	a1 = 1
	fmt.Println(message, a1, b1, c1)
	// once
	var a2, b2, c2 int = 1, 2, 3
	fmt.Println(message, a2, b2, c2)

	// 同类型定义时可省略类型
	var a3, b3, c3 = 1, 2, 3
	fmt.Println(a3, b3, c3)

	var a4, b4, c4 = 1, false, 3
	fmt.Println(a4, b4, c4)

	// 短赋值
	message2 := "hello2"
	a5, b5, c5 := 1, false, 3
	fmt.Println(message2, a5, b5, c5)

	fmt.Println("all together")
	d = "ddddd"
	var e = 42
	f := 43
	g := "gggg"
	h, i := "stored h", "stored i"
	j, k, l, m := 44.7, true, false, 'm'
	n := "n"
	o := `o`
	fmt.Println("a - ", a)
	fmt.Println("b - ", b)
	fmt.Println("c - ", c)
	fmt.Println("d - ", d)
	fmt.Println("e - ", e)
	fmt.Println("f - ", f)
	fmt.Println("g - ", g)
	fmt.Println("h - ", h)
	fmt.Println("i - ", i)
	fmt.Println("j - ", j)
	fmt.Println("k - ", k)
	fmt.Println("l - ", l)
	fmt.Println("m - ", m)
	fmt.Println("n - ", n)
	fmt.Println("o - ", o)

	fmt.Println("----------")
	var name1 = "Todd1"
	name2 := "name2"
	name3 := `name3`
	fmt.Println("hello, ", name)
	fmt.Println("hello, ", name1)
	fmt.Println("hello, ", name2)
	fmt.Println("hello, ", name3)
}
