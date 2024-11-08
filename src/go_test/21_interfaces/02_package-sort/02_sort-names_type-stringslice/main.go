package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

func (p person) String() string {
	return fmt.Sprintf("YAYAYA %s: %d", p.Name, p.Age)
}

type ByAge []person

func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func main() {
	s := []string{"Zero", "John", "Al", "Jenny"}
	fmt.Println(s)
	//sort.StringSlice(s).Sort()
	//sort.Sort(sort.StringSlice(s))
	//sort.Strings(s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)

	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println(n)
	//sort.Sort(sort.IntSlice(n))
	//sort.Sort(sort.Reverse(sort.IntSlice(n)))
	sort.Ints(n)
	fmt.Println(n)

	fmt.Println("====================")
	people := []person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jennny", 26},
	}
	fmt.Println(people[0])
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
}
