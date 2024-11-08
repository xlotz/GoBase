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
	res.Body.Close()
	sc.Split(bufio.ScanWords)

	buck := make([]int, 12)
	for sc.Scan() {
		n := hashBucket3(sc.Text(), 12)
		buck[n]++
	}
	fmt.Println(buck)

}

func hashBucket3(w string, b int) int {
	var sum int
	for _, v := range w {
		sum += int(v)
	}
	return sum % b
}
