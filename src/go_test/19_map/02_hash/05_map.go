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

	bu := make(map[int]map[string]int)
	for i := 0; i < 12; i++ {
		bu[i] = make(map[string]int)
	}

	for sc.Scan() {
		w := sc.Text()
		n := hashBucket5(w, 12)
		bu[n][w]++
	}
	for k, v := range bu[6] {
		fmt.Println(v, "\t- ", k)
	}
}

func hashBucket5(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
