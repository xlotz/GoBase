package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {

	res, err := http.Get("https://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	sc := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	sc.Split(bufio.ScanWords)
	buck := make([]int, 200)
	for sc.Scan() {
		n := hashBucket1(sc.Text())
		buck[n]++
	}
	fmt.Println(buck[65:123])
}

func hashBucket1(w string) int {
	return int(w[0])
}
