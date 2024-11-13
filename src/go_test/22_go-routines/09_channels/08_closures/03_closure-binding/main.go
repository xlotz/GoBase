package main

import "fmt"

func main() {
	done := make(chan bool)

	va := []string{"a", "b", "c"}
	for _, v := range va {
		v1 := v
		go func() {
			fmt.Println(v1)
			done <- true
		}()
	}

	for _ = range va {
		<-done
	}
}
