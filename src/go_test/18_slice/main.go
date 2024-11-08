package main

import (
	"fmt"
)

func main() {

	// int
	fmt.Println("int slice ...")
	s1 := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("%T\n", s1)
	fmt.Println(s1)

	for i, v := range s1 {
		fmt.Println(i, " - ", v)
	}

	s2 := make([]int, 0, 3)
	fmt.Println("---------------")
	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
	fmt.Println("---------------")
	for i := 0; i < 80; i++ {
		s2 = append(s2, i)
		fmt.Println("len: ", len(s2), "cap: ", cap(s2), "v: ", s2[i])
	}

	fmt.Println("string slice ...")
	greet1 := []string{
		"Good morning!",
		"Bonjour!",
		"dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	for i, c := range greet1 {
		fmt.Println(i, c)
	}
	for j := 0; j < len(greet1); j++ {
		fmt.Println(greet1[j])
	}

	fmt.Println("-------------------")
	var res []int
	fmt.Println(res)
	s3 := []string{"a", "b", "c", "e", "f", "g"}
	fmt.Println(s3)
	fmt.Println(s3[2:4])
	fmt.Println(s3[2])
	fmt.Println("mystring"[2])
	fmt.Println("-------------------")
	fmt.Print("[1:2] ")
	fmt.Println(greet1[1:2])
	fmt.Print("[:2] ")
	fmt.Println(greet1[:2])
	fmt.Print("[5:] ")
	fmt.Println(greet1[5:])
	fmt.Print("[:] ")
	fmt.Println(greet1[:])

	fmt.Println("make ...")
	c1 := make([]int, 3)
	c1[0] = 7
	c1[1] = 10
	c1[2] = 15
	fmt.Println(c1[0], c1[1], c1[2])

	greet2 := make([]string, 3, 5)
	greet2[0] = "Good morning"
	greet2[1] = "Bonjour"
	greet2[2] = "dias"

	fmt.Println(greet2[2])
	//greet2[3] = "supr"  // 超出长度， 报错
	//fmt.Println(greet2[2])

	greet2 = append(greet2, "supraaa")
	fmt.Println(greet2[3])

	greet2 = append(greet2, "1")
	greet2 = append(greet2, "2")
	greet2 = append(greet2, "3")
	greet2 = append(greet2, "4")
	fmt.Println(greet2[6])
	fmt.Println(len(greet2), cap(greet2))

	fmt.Println("s1 ...")
	s4 := []int{6, 7, 8, 9}
	s1 = append(s1, s4...)
	fmt.Println(s1)

	fmt.Println("----------------")
	s5 := []string{"aa", "bb"}
	s6 := []string{"cc", "dd"}
	s5 = append(s5, s6...)
	fmt.Println(s5)
	s5 = append(s5[:2], s5[3:]...)
	fmt.Println(s5)

	fmt.Println("short...")
	st := []string{}
	sts := [][]string{}
	fmt.Println(st, sts, st == nil)

	var st1 []string
	var sts1 [][]string
	fmt.Println(st1, sts1, st1 == nil) //空值

	st2 := make([]string, 35)
	sts2 := make([][]string, 35)
	fmt.Println(st2, sts2, st2 == nil) //非空值

	fmt.Println("---------------------")
	st3 := []string{}
	sts3 := [][]string{}
	//st3[0] = "Todd" // 超长，报错
	//st3 = append(st3, "bbb")
	fmt.Println(st3, sts3)

	var st4 []string
	var sts4 [][]string
	//st4[0] = "Todd" // 超长报错
	//st4 = append(st4, "AAAA")
	fmt.Println(st4, sts4)

	fmt.Println("---------------------")
	st5 := make([]string, 35)
	sts5 := make([][]string, 35)
	st5[0] = "Todd"
	st5 = append(st5, "AAAA")
	fmt.Println(st5, sts5)

	fmt.Println("---------------------")

	var rec [][]string
	st6 := make([]string, 4)
	st6[0] = "A"
	st6[1] = "B"
	st6[2] = "C"
	st6[3] = "D"
	rec = append(rec, st6)
	st7 := make([]string, 4)
	st7[0] = "Aa"
	st7[1] = "Ba"
	st7[2] = "Ca"
	st7[3] = "Da"
	rec = append(rec, st7)
	fmt.Println(rec)

	fmt.Println("---------------------")
	trs := make([][]int, 0, 3)
	for i := 0; i < 3; i++ {
		tr := make([]int, 0, 4)
		for j := 0; j < 4; j++ {
			tr = append(tr, j)
		}
		trs = append(trs, tr)
	}
	fmt.Println(trs)

	fmt.Println("---------------------")
	m1 := make([]int, 1)
	fmt.Println(m1[0])
	m1[0] = 7
	fmt.Println(m1[0])
	m1[0]++
	fmt.Println(m1[0])

}
