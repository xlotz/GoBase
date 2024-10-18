package main

import (
	//"HelloWorld/src/17_Iter/iterx"
	"bufio"
	"fmt"
	"io"
	"iter"
	"maps"
	"os"
	"slices"
)

func Fibonacci1(n int) func() (int, bool) {
	a, b, c := 1, 1, 2
	i := 0
	return func() (int, bool) {
		if i >= n {
			return 0, false
		} else if i < 2 {
			f := i
			i++
			return f, true
		}
		a, b = b, c
		c = a + b
		i++
		return a, true
	}
}

func Fib(n int) func() (int, bool) {
	a, b, c := 1, 1, 2
	i := 0
	return func() (int, bool) {
		if i >= n {
			return 0, false
		} else if i < 2 {
			f := i
			i++
			return f, true
		}

		a, b = b, c
		c = a + b
		i++

		return a, true
	}
}

func Fibonacci2(n int) func(yield func(int) bool) {
	a, b, c := 0, 1, 1
	return func(yield func(int) bool) {
		for range n {
			if !yield(a) {
				return
			}
			a, b = b, c
			c = a + b
		}
	}
}
func main() {
	n := 8
	for f := range Fibonacci2(n) {
		fmt.Println(f)
	}

	// 拉取式迭代器
	next, stop := iter.Pull(Fibonacci2(n))
	defer stop()
	for {
		fibn, ok := next()
		if !ok {
			break
		}
		fmt.Println(fibn)
	}

	// 推送式迭代器，错误处理演示
	fmt.Println("推送式迭代器，错误处理演示")
	file, _ := os.Open("README.md")
	for line, err := range ScanLines(file) {
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(line)
	}

	// 拉取式迭代器，错误处理演示
	fmt.Println("拉取式迭代器，错误处理演示")
	next2, stop := iter.Pull2(ScanLines(file))
	defer stop()
	for {
		line, err, ok := next2()
		if err != nil {
			fmt.Println(err)
			break
		} else if !ok {
			break
		}
		fmt.Println(line)
	}

	//
	fmt.Println("标准库演示")
	fmt.Println()

	fmt.Println("slice")
	s := []int{1, 2, 3, 4, 5}
	for i, n := range slices.All(s) {
		fmt.Println(i, n)
	}

	for n := range slices.Values(s) {
		fmt.Println(n)
	}
	for chunk := range slices.Chunk(s, 2) {
		fmt.Println(chunk)
	}

	s2 := slices.Collect(slices.Values(s))
	fmt.Println(s2)

	fmt.Println()
	fmt.Println("map ...")
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	keys := slices.Collect(maps.Keys(m))
	fmt.Println(keys)

	value := slices.Collect(maps.Values(m))
	fmt.Println(value)

	for k, v := range maps.All(m) {
		fmt.Println(k, v)
	}

	m2 := maps.Collect(maps.All(m))
	fmt.Println(m2)

	//// iterx
	//s := []string{"apple", "banana", "cherry"}
	//all := iterx.Slice(s).Map(strings.ToUpper).All()
	//for i, v := range all {
	//	fmt.Println("index: %d, value: %s\n", i, v)
	//}
}

func ScanLines(reader io.Reader) iter.Seq2[string, error] {
	scanner := bufio.NewScanner(reader)
	return func(yield func(string, error) bool) {
		for scanner.Scan() {
			if !yield(scanner.Text(), scanner.Err()) {
				return
			}
		}
	}
}
