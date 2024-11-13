package main

import "fmt"

func main() {
	x := 1
	str := evalInt1(x)
	fmt.Println(str)

	str2 := evalInt2(x)
	fmt.Println(str2)
}

func evalInt1(n int) string {
	if n > 10 {
		return fmt.Sprintf("x is greater than 10")
	} else {
		return fmt.Sprintf("x is less than 10")
	}
}

func evalInt2(n int) string {
	if n > 10 {
		return fmt.Sprintf("x is greater than 10")
	}
	return fmt.Sprintf("x is less than 10")

}
