package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	mp1 := map[int]string{
		0: "a",
		1: "b",
		2: "c",
		3: "d",
		4: "e",
	}
	mp2 := map[string]int{
		"a": 0,
		"b": 1,
		"c": 2,
		"d": 3,
		"e": 4,
	}
	mp3 := make(map[string]int, 8)
	mp4 := make(map[string][]int, 10)

	fmt.Println(mp1[0])
	fmt.Println(mp2["f"])
	fmt.Println(len(mp3))
	mp4["a"] = append(mp4["a"], 1, 2, 3)
	fmt.Println(mp4["a"])

	if _, exist := mp2["b"]; exist {
		mp2["b"] = 3
	}
	fmt.Println(mp2["b"])

	// NaN
	mp5 := make(map[float64]string, 10)
	mp5[math.NaN()] = "a"
	mp5[math.NaN()] = "b"
	mp5[math.NaN()] = "c"
	_, exist := mp5[math.NaN()]
	fmt.Println(exist)
	fmt.Println(mp5)

	// delete

	fmt.Println(mp2)
	delete(mp2, "a")
	fmt.Println(mp2)

	delete(mp5, math.NaN())
	fmt.Println(mp5)

	// for range
	for k, v := range mp2 {
		fmt.Println(k, v)
	}

	for k, v := range mp5 {
		fmt.Println(k, v)
	}

	// clear

	clear(mp2)
	fmt.Println(mp2)

	// set

	set := make(map[int]struct{}, 10)
	for i := 0; i < 10; i++ {
		set[rand.Intn(100)] = struct{}{}
	}
	fmt.Println(set)
}
