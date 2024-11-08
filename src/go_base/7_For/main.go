package main

import "fmt"

func main() {
	// for
	fmt.Println("for")
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	for i, j := 1, 2; i < 10 && j < 100; i, j = i+1, j+10 {
		fmt.Println(i, j)
	}

	fmt.Println("while")
	num := 1
	for num < 100 {
		num *= 2
		fmt.Println(num)
	}

	// 9*9
	fmt.Println("9 * 9")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if i <= j {
				fmt.Printf("%d * %d = %2d ", i, j, i*j)
			}
		}
		fmt.Println()
	}

	// for range
	seq := "hello world"
	for i, v := range seq {
		fmt.Println(i, v)
	}

	// break
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i <= j {
				break
			}
			fmt.Println(i, j)
		}
	}
	fmt.Println("break Out")
Outer:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i <= j {
				break Outer
			}
			fmt.Println(i, j)
		}
	}

	fmt.Println("continue")
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i > j {
				continue
			}
			fmt.Println(i, j)
		}
	}

	fmt.Println("continue Out")
Out:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i > j {
				continue Out
			}
			fmt.Println(i, j)
		}
	}

}
