package main

import "fmt"

func main() {
	// non-interface-error

	//name1 := "Sydney"
	//str, ok := name1.(string)
	//if ok {
	//	fmt.Printf("%q\n", str)
	//} else {
	//	fmt.Printf("v is not a string\n")
	//}

	// interface string
	var name2 interface{} = "Sydney"
	str2, ok2 := name2.(string)
	if ok2 {
		fmt.Printf("name2 %T\n", str2)
	} else {
		fmt.Printf("v is not a string\n")
	}

	// interface string not ok
	var name3 interface{} = 7
	str3, ok3 := name3.(string)
	if ok3 {
		fmt.Printf("name3 %T\n", str3)
	} else {
		fmt.Printf("v is not a string\n")
	}

	fmt.Println("interface int sum...")
	var val interface{} = 7
	fmt.Printf("%T\n", val)
	//fmt.Println(val + 6) //invalid operation: val + 6 (mismatched types interface{} and int)

	fmt.Println(val.(int) + 6)

	fmt.Println("casting reminder ...")
	rem := 7.24
	fmt.Printf("%T\n", rem)
	fmt.Printf("%T\n", int(rem))

	fmt.Println("interface cast error ...")
	var val1 interface{} = 7
	fmt.Printf("%T\n", val1)
	//fmt.Printf("%T\n", int(val1)) // interface cast error
	fmt.Printf("%T\n", val1.(int))
}
