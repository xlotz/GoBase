package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 输出
	//os.Stdout.WriteString("hello 1")
	//fmt.Println()
	//print("hello 2\n")
	//println("hello 3")
	//fmt.Println()
	//fmt.Println("hello 4")
	//
	//writer := bufio.NewWriter(os.Stdout)
	//defer writer.Flush()
	//writer.WriteString("hello 5")
	//fmt.Println(writer, "hello  6")
	//fmt.Println()
	//fmt.Printf("hello 7 %s!", "jack")
	//fmt.Println()
	// 输入

	//var buf [1024]byte
	//n, _ := os.Stdin.Read(buf[:])
	//os.Stdout.Write(buf[:n])

	//var a, b int
	//fmt.Scanln(&a, &b)
	//fmt.Printf("%d + %d = %d", a, b, a+b)

	//n := 10
	//s := make([]int, n)
	//for _, i := range n {
	//	fmt.Scan(&s[i])
	//}
	//fmt.Println(s)

	//reader := bufio.NewReader(os.Stdin)
	//var a, b int
	//fmt.Scanln(reader, &a, &b)
	//fmt.Printf("%d + %d = %d\n", a, b, a+b)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "exit" {
			break
		}
		fmt.Println("scan", line)

	}

}
