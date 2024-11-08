package main

import (
	"encoding/json"
	"fmt"
)

type person1 struct {
	First string
	Last  string
	Age   int
}

type person2 struct {
	First string
	Last  string
	Age   int `json:"wisdom score"`
}

func main() {
	var p1 person1
	fmt.Println(p1.First, p1.Last, p1.Age)

	bs1 := []byte(`{"First":"A", "Last":"B", "Age":20}`)
	json.Unmarshal(bs1, &p1)
	fmt.Println("----------------")
	fmt.Println(p1.First, p1.Last, p1.Age)
	fmt.Printf("%T \n", p1)

	fmt.Println("====================")
	var p2 person2
	fmt.Println(p2.First, p2.Last, p2.Age)

	bs2 := []byte(`{"First":"A", "Last":"B", "wisdom score":20}`)
	json.Unmarshal(bs2, &p2)
	fmt.Println("----------------")
	fmt.Println(p2.First, p2.Last, p2.Age)
	fmt.Printf("%T \n", p2)
}
