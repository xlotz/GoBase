package main

import (
	"encoding/json"
	"fmt"
)

type person1 struct {
	First       string
	Last        string
	Age         int
	notExported int
}

type person2 struct {
	first string
	last  string
	age   int
}

type person3 struct {
	First string
	Last  string `json:"-"`
	Age   int    `json:"wisdom score"`
}

func main() {
	p1 := person1{"Tom", "Bond", 20, 007}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))

	fmt.Println()

	p2 := person2{"Jack", "AAA", 30}
	fmt.Println(p2)
	bs2, _ := json.Marshal(p2)
	fmt.Println(string(bs2))

	fmt.Println()
	p3 := person3{"AAA", "BBB", 20}
	bs3, _ := json.Marshal(p3)
	fmt.Println(string(bs3))
}
