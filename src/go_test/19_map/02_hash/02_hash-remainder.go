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
	bu := make([]int, 12)
	for sc.Scan() {
		n := hashBucket2(sc.Text(), 12)
		bu[n]++
	}
	fmt.Println(bu)
}

func hashBucket2(word string, buckets int) int {
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}
