package main

import "fmt"

type Person interface {
	Say(string) string
	Walk(int)
}

type Man interface {
	Exercise()
	Person
}

// 空接口
type Any interface {
}

// 起重机接口
type Crane interface {
	JackUp() string
	Hoist() string
}

// A公司
type CraneA struct {
	work int //内部字段不同，代表内部细节不同
}

func (c CraneA) Work() {
	fmt.Println("使用A公司技术")
}
func (c CraneA) JackUp() string {
	c.Work()
	return "jackUP"
}
func (c CraneA) Hoist() string {
	c.Work()
	return "Hoist"
}

// B公司

type CraneB struct {
	work int //内部字段不同，代表内部细节不同
}

func (c CraneB) Work() {
	fmt.Println("使用B公司技术")
}
func (c CraneB) JackUp() string {
	c.Work()
	return "jackUP"
}
func (c CraneB) Hoist() string {
	c.Work()
	return "Hoist"
}

type Company struct {
	Crane Crane // 根据公司类型存放起重机
}

func (c *Company) Build() {
	fmt.Println(c.Crane.JackUp())
	fmt.Println(c.Crane.Hoist())
	fmt.Println("建设完成")
}
func main() {
	// 仅仅只有接口无法初始化，但可声明
	var person Person
	fmt.Println(person)

	comp := Company{CraneA{}}
	comp.Build()
	fmt.Println()

	comp.Crane = CraneB{}
	comp.Build()

	//空接口
	var any Any
	any = 1
	println(any)
	fmt.Println(any)
	any = []int{}
	println(any)
	fmt.Println(any)

	any = map[string]int{}
	println(any)
	fmt.Println(any)
}
