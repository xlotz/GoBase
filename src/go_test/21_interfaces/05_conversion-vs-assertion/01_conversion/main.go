package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x, y = 12, 12.1230123
	fmt.Println(y + float64(x))
	fmt.Println(int(y) + x)

	var a rune = 'a'
	var b int32 = 'b'
	fmt.Println(a, b, string(a), string(b))

	fmt.Println(string([]byte{'h', 'e', 'l', 'l', 'o'}))
	fmt.Println([]byte("hello"))

	fmt.Println("strconv...")
	fmt.Println("Atoi...")
	var x1 = "12"
	var y1 = 6
	z1, _ := strconv.Atoi(x1)
	fmt.Println(y1 + z1)
	fmt.Println("ItoA...")
	x2 := 12
	y2 := "I have this many: " + strconv.Itoa(x2)
	fmt.Println(y2)

	fmt.Println("------------")
	b1, _ := strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("3.1415", 64)
	i, _ := strconv.ParseInt("-42", 10, 64)
	u, _ := strconv.ParseUint("42", 10, 64)
	println(b1, f, i, u)

	w := strconv.FormatBool(true)
	x3 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	y3 := strconv.FormatInt(-42, 16)
	z3 := strconv.FormatUint(42, 16)
	fmt.Println(w, x3, y3, z3)
}
