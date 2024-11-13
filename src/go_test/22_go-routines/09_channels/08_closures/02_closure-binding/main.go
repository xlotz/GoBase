package main

import "fmt"

func main() {
	done := make(chan bool)
	va := []string{"a", "b", "c"}
	for _, v := range va {
		go func(u string) {
			fmt.Println(u)
			done <- true
		}(v)
	}

	for _ = range va {
		<-done
	}
}
