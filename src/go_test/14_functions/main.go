package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	greet("jone")
	greet1("s", "x")
	fmt.Println(greet2("aaa", "Bbb"))
	fmt.Println(greet3("bbba", "Bbb"))
	fmt.Println(greet4("bbba", "Bbb"))

	n1 := average(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(n1)
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n2 := average(data...)
	fmt.Println(n2)
	n3 := average2(data)
	fmt.Println(n3)

	greet5 := makeGreeter()
	fmt.Println(greet5())
	//
	fmt.Printf("%T\n", greet5)

	//
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "The xxx"
		fmt.Println(y)
	}

	//
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2

	x2 := 0
	incr := func() int {
		x2++
		return x2
	}
	fmt.Println(incr()) // 1
	fmt.Println(incr()) // 2

	incr2 := wrapper()
	fmt.Println(incr2()) // 1
	fmt.Println(incr2()) // 2

	incrA := wrapper()
	incrB := wrapper()
	fmt.Println("A:", incrA())
	fmt.Println("A:", incrA())
	fmt.Println("B:", incrB())
	fmt.Println("B:", incrB())
	fmt.Println("B:", incrB())

	// callback print nums
	fmt.Println("callback ...")
	visit([]int{1, 2, 3, 4}, func(i int) {
		fmt.Println(i)
	})

	fmt.Println("callback filter")
	xs := filter([]int{1, 2, 3, 4}, func(i int) bool {
		return i > 1
	})
	fmt.Println(xs)
	fmt.Println("递归...")
	fmt.Println(factorial(4))

	fmt.Println("延迟调用...")
	world()
	hello()

	fmt.Println("----------")
	defer world()
	hello()

	fmt.Println("int ...")
	age := 44
	changeMe(age)
	fmt.Println("main: ", age)
	fmt.Println("int end")
	fmt.Println()

	fmt.Println("int pointer ...")
	fmt.Println("main: ", &age)
	changeMe2(&age)
	fmt.Println("main: ", &age)
	fmt.Println("main: ", age)
	fmt.Println("int pointer end")
	fmt.Println()

	fmt.Println("string ... ")
	name := "Tom"
	fmt.Println("main: ", name)
	changeMe3(name)
	fmt.Println("main: ", name)

	fmt.Println("string pointer ... ")
	name2 := "Todd"
	fmt.Println("main: ", name2)
	changeMe4(&name2)
	fmt.Println("main: ", name2)

	fmt.Println()
	fmt.Println("reference-type")
	m1 := make(map[string]int)
	changeMe5(m1)
	fmt.Println(m1["ddd"])

	m2 := make([]string, 1, 25)
	fmt.Println("main: ", m2)
	changeMe6(m2)
	fmt.Println("main: ", m2)

	fmt.Println()
	c1 := customer{"Tome", 44}
	fmt.Println("main: ", &c1.name)
	changeMe7(&c1)
	fmt.Println("main: ", c1)
	fmt.Println("main: ", &c1.name)

	//匿名函数
	func() {
		fmt.Println("I'm driving!")
	}()
}

func greet(name string) {
	fmt.Println(name)
}

func greet1(fname string, lname string) {
	fmt.Println(fname, lname)
}

func greet2(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}

func greet3(fname, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}

func greet4(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

func average2(sf []float64) float64 {
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

/*
*
 */
func makeGreeter() func() string {
	return func() string {
		return "Hello world"
	}
}

/*
*
 */
var x1 int // 零值
func increment() int {
	x1++
	return x1
}

/*
*
 */
func wrapper() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

/**
print nums
调用其他函数
*/

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

/**
filter nums
调用其他函数
*/

func filter(numbers []int, callback func(int) bool) []int {
	var xs []int
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	return xs
}

/**
递归
*/

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

/*
*
defer
*/
func hello() {
	fmt.Print("hello 1")
}
func world() {
	fmt.Print("world 2")
}

/**

 */

func changeMe(z int) {
	fmt.Println("修改前: ", z)
	z = 24
	fmt.Println("修改后: ", z)
}

func changeMe2(z *int) {
	fmt.Println("修改前: ", z)  // address
	fmt.Println("修改前: ", *z) // 44
	*z = 24
	fmt.Println("修改后: ", z)  // address
	fmt.Println("修改后: ", *z) // 24
}

func changeMe3(z string) {
	fmt.Println("修改前: ", z)
	z = "Rock"
	fmt.Println("修改后: ", z)
}

func changeMe4(z *string) {
	fmt.Println("修改前: ", z)
	fmt.Println("修改前: ", *z)
	*z = "RRRR"
	fmt.Println("修改后: ", z)
	fmt.Println("修改后: ", *z)
}

/**
map
*/

func changeMe5(z map[string]int) {
	z["ddd"] = 44
}

func changeMe6(z []string) {
	fmt.Println("修改前: ", z[0])
	z[0] = "aaa"
	fmt.Println("修改后: ", z[0])

}

type customer struct {
	name string
	age  int
}

func changeMe7(z *customer) {
	fmt.Println("修改前: ", z)
	fmt.Println("修改前: ", &z.name)
	z.name = "DDD"
	fmt.Println("修改后: ", z)
	fmt.Println("修改后: ", &z.name)
}
