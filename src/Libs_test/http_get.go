package main

import (
	"fmt"
	"io"
	http2 "net/http"
)

func main() {
	http, err := http2.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer http.Body.Close()
	cont, err := io.ReadAll(http.Body)
	fmt.Println(cont)
}
