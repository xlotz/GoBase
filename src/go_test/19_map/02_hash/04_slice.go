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
	bu := make([][]string, 12)

	for sc.Scan() {
		w := sc.Text()
		n := hashBucket4(w, 12)
		bu[n] = append(bu[n], w)
	}

	for i := 0; i < 12; i++ {
		fmt.Println(i, " - ", len(bu[i]))
	}
	fmt.Println(len(bu), cap(bu))
}

func hashBucket4(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
	// comment out the above, then uncomment the below
	// a more uneven distribution
	// return len(word) % buckets
}
