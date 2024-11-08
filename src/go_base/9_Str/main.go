package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	str := "this is a string"
	fmt.Println(str[0])
	fmt.Println(string(str[0]))

	bytes := []byte(str)
	fmt.Println(bytes)
	fmt.Println(string(bytes))

	fmt.Println(&str)
	bytes = append(bytes, 96, 95, 97, 98, 99)
	fmt.Println(string(bytes))

	str2 := "这是一个字符串"
	fmt.Println(len(str), len(str2))
	fmt.Println(string(str[0]), string(str2[0]), string(str2[0:3]))

	fmt.Println("拷贝")
	var dst, src, dst2 string
	src = str
	desBytes := make([]byte, len(src))
	copy(desBytes, src)
	dst = string(desBytes)
	fmt.Println(src, dst)

	dst2 = strings.Clone(src)
	fmt.Println(src, dst2)

	fmt.Println("拼接")
	str = str + " aaaa"
	fmt.Println(str)
	byte2 := []byte(str)
	byte2 = append(byte2, " append"...)
	fmt.Println(string(byte2))

	fmt.Println("拼接 strings.Builder")
	builder := strings.Builder{}
	builder.WriteString(str)
	builder.WriteString("buildersss")
	fmt.Println(builder.String())

	fmt.Println("遍历")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%d, %x, %s\n", str[i], str[i], string(str[i]))
	}
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%d, %x, %s\n", str2[i], str2[i], string(str2[i]))
	}
	for _, r := range str2 {
		fmt.Printf("%d, %x, %s\n", r, r, string(r))
	}
	fmt.Println("runes")
	runes := []rune(str2)
	for i := 0; i < len(runes); i++ {
		fmt.Println(string(runes[i]))
	}

	fmt.Println("utf8")

	for i, w := 0, 0; i < len(str2); i += w {
		r, width := utf8.DecodeRuneInString(str2[i:])
		fmt.Println(string(r))
		w = width
	}
}
