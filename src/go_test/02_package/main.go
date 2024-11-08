package main

import (
	"fmt"
	"go_test/02_package/icome"
	"go_test/02_package/stringutil"
)

func main() {

	fmt.Println(stringutil.Reverse("Hello, go!"))
	fmt.Println(stringutil.MyName)
	fmt.Println(icome.BearName)

}
