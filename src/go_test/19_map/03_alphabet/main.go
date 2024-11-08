package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	bs, _ := ioutil.ReadAll(res.Body)
	str := string(bs)
	fmt.Println(str)
	fmt.Println("-----------------")
	w := make(map[string]string)
	sc := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		w[sc.Text()] = ""
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "input: ", err)
	}
	i := 0
	for k := range w {
		fmt.Println(k)
		if i == 200 {
			break
		}
		i++
	}

}
