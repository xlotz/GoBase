package main

import "fmt"

func main() {

	//
	if true {
		fmt.Println("This run")
	}
	if false {
		fmt.Println("This did not run")
	}

	// not
	if !true {
		fmt.Println("This did not run")
	}
	if !false {
		fmt.Println("This run")
	}

	// or
	if true || false {
		fmt.Println("This run")
	}

	// and

	if true && false {
		fmt.Println("This did not run")
	}

}
