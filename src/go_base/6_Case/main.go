package main

import "fmt"

func main() {
	//a, b := 1, 2

	//if a > b {
	//	b++
	//	fmt.Println(b)
	//} else {
	//	a++
	//	fmt.Println(a)
	//}

	//if a<<1%100+3 > b*100/20+6 {
	//	b++
	//	fmt.Println(b)
	//
	//} else {
	//	a++
	//	fmt.Println(a)
	//}

	//if x := 0 + 1; x > 2 {
	//	fmt.Println(x)
	//}
	// 从左到右
	//score := 90
	//var ans string
	//if score == 100 {
	//	ans = "s"
	//} else if score >= 90 && score < 100 {
	//	ans = "a"
	//} else if score >= 80 && score < 90 {
	//	ans = "b"
	//} else if score >= 70 && score < 80 {
	//	ans = "c"
	//} else if score >= 60 && score < 70 {
	//	ans = "d"
	//} else if score >= 50 && score < 60 {
	//	ans = "e"
	//} else if score >= 40 && score < 50 {
	//	ans = "f"
	//} else {
	//	ans = "nil"
	//}
	//fmt.Println(ans)

	// 从上到下

	//if score >= 0 && score < 60 {
	//	ans = "f"
	//} else if score >= 0 && score < 60 {
	//	ans = "e"
	//} else if score < 70 {
	//	ans = "d"
	//} else if score < 80 {
	//	ans = "c"
	//} else if score < 90 {
	//	ans = "b"
	//} else if score < 100 {
	//	ans = "a"
	//} else if score == 100 {
	//	ans = "s"
	//} else {
	//	ans = "nil"
	//}
	//fmt.Println(ans)

	// switch
	//str := "a"
	//switch str {
	//case "a":
	//	str += "a"
	//	str += "c"
	//case "b":
	//	str += "bb"
	//	str += "aaaa"
	//default:
	//	str += "cccc"
	//
	//}
	//fmt.Println(str)

	switch num := f(); {
	case num >= 0 && num <= 1:
		num++
		fmt.Println("----", num)
	case num > 1:
		num--
		fmt.Println(num)
		fallthrough

	case num < 0:
		num += num
		fmt.Println(num)
	}

	a := 1
	if a == 1 {
		goto A
	} else {
		fmt.Println("b")
	}
A:
	fmt.Println("a")

}

func f() int {
	return 1
}
