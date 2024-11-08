package main

import (
	"flag"
	"fmt"
)

func main() {

	//name := flag.String("name", "张三", "姓名")
	//age := flag.Int("age", 15, "年龄")
	//sex := flag.Bool("sex", true, "性别")

	var name string
	var age int
	var sex bool
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 15, "年龄")
	flag.BoolVar(&sex, "sex", true, "性别")

	flag.Parse()
	fmt.Println(name, age, sex)

}
