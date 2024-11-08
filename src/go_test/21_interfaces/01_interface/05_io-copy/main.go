package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	msg := "Do not dwell in the past, do not dream of the future|"
	rdr := strings.NewReader(msg)
	io.Copy(os.Stdout, rdr)

	rdr2 := bytes.NewBuffer([]byte(msg))
	io.Copy(os.Stdout, rdr2)

	res, _ := http.Get("https://www.baidu.com")
	io.Copy(os.Stdout, res.Body)
	res.Body.Close()

	fmt.Println("----------------------------")

	msg2 := "concentrate the mind on the present"
	rd1 := strings.NewReader(msg2)
	_, err := io.Copy(os.Stdout, rd1)
	if err != nil {
		fmt.Println(err)
		return
	}

	rd2 := bytes.NewBuffer([]byte(msg2))
	_, err1 := io.Copy(os.Stdout, rd2)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	res2, err2 := http.Get("https://www.baidu.com")
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	io.Copy(os.Stdout, res2.Body)
	if err3 := res2.Body.Close(); err3 != nil {
		fmt.Println(err3)
	}
}
