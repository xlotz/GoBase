package main

import "fmt"

func main() {
	// int
	fmt.Println("int loop...")
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	// nested
	fmt.Println("nested ...")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, " - ", j)
		}
	}
	// for condition while
	fmt.Println("for condition while ...")
	i1 := 0
	for i1 < 10 {
		fmt.Println(i1)
		i1++
	}

	fmt.Println("for no condition ...")
	//for {
	//	fmt.Println(i1)
	//	i1++
	//}

	fmt.Println("for break ...")
	for {
		fmt.Println(i1)
		if i1 >= 15 {
			break
		}
		i1++

	}

	fmt.Println("for continue ...")
	i2 := 0
	for {
		i2++
		if i2%2 == 0 {
			continue
		}
		fmt.Println(i2)
		if i2 >= 50 {
			break
		}
	}

	fmt.Println("rune loop ...")
	for i := 250; i < 340; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
	foo2 := "a"
	fmt.Println(foo2)
	fmt.Printf("%T \n", foo2)

	fmt.Println("rune loop ...")
	for i := 50; i <= 100; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}

}
