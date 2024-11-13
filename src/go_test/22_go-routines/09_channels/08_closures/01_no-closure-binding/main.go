package main

import (
	"fmt"
)

func main() {
	done := make(chan bool)
	va := []string{"a", "b", "c"}
	for _, v := range va {
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	for _ = range va {
		<-done
	}
}
