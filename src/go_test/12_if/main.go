package main

import "fmt"

func main() {
	// eval true
	fmt.Println("eval true ...")
	if true {
		fmt.Println("This ran")
	}
	if false {
		fmt.Println("This did not run")
	}

	// not exclamation
	fmt.Println("not exclamation ...")
	if !true {
		fmt.Println("This did not run")
	}
	if !false {
		fmt.Println("This ran")
	}

	// init statement
	fmt.Println("init statement ...")
	b := true
	if food := "Chocolate"; b {
		fmt.Println(food)
	}

	// init statement error invalid
	fmt.Println("init statement error invalid ...")
	if food := "Chocolate"; b {
		fmt.Println(food)
	}
	//fmt.Println(food) // 没有定义，编译出错

	// if else
	fmt.Println("if else ...")
	if false {
		fmt.Println("first print statement")
	} else {
		fmt.Println("second print statement")
	}

	// if elif else
	fmt.Println("if else if else ...")
	if false {
		fmt.Println("first print statement")
	} else if true {
		fmt.Println("second print statement")
	} else {
		fmt.Println("third print statement")
	}

	if false {
		fmt.Println("first print statement")
	} else if false {
		fmt.Println("second print statement")
	} else if true {
		fmt.Println("ahaaa print statement")
	} else {
		fmt.Println("third print statement")
	}

	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
