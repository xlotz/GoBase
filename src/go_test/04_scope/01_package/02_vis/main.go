package main

import (
	"fmt"
	"go_test/04_scope/01_package/02_vis/vis"
)

func main() {

	fmt.Println(vis.MyName)
	vis.PrintVar()

}
