package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	fmt.Println("nil map...")
	fmt.Println("定义...")

	var m1 map[string]string
	fmt.Println(m1, m1 == nil)

	var m2 = make(map[string]string)
	m2["a"] = "good"
	m2["b"] = "bon"
	fmt.Println(m2)

	m3 := make(map[string]string)
	m3["a"] = "good"
	m3["b"] = "bon"
	fmt.Println(m3)

	m4 := map[string]string{}
	m4["a"] = "good"
	m4["b"] = "bon"
	fmt.Println(m4)

	m5 := map[string]string{
		"a": "good",
		"b": "bon",
	}
	fmt.Println(m5)
	fmt.Println("添加...")
	m5["c"] = "aaa"
	fmt.Println(m5)

	fmt.Println("长度...")
	fmt.Println(len(m5))

	fmt.Println("更新...")
	m5["a"] = "aaa"
	m5["b"] = "bbb"
	fmt.Println(m5)

	fmt.Println("删除...")
	delete(m5, "a")
	fmt.Println(m5)

	fmt.Println("判断是否存在...")
	m6 := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}
	fmt.Println(m6)
	if v, exists := m6[2]; exists {
		fmt.Println("v: ", v, " , exists: ", exists)
	} else {
		fmt.Println("v: ", v, " , exists: ", exists)
	}
	fmt.Println(m6)

	fmt.Println("删除不存在的key...")
	delete(m6, 7)
	fmt.Println(m6)

	if v, exists := m6[2]; exists {
		delete(m6, 2)
		fmt.Println("v: ", v, " , exists: ", exists)
	} else {
		fmt.Println("v: ", v, " , exists: ", exists)
	}
	fmt.Println(m6)
	fmt.Println("迭代...")
	for k, v := range m6 {
		fmt.Println(k, " - ", v)
	}

	fmt.Println("hash table ...")
	le := 'A'
	fmt.Println(le)
	fmt.Printf("%T \n", le)
	le1 := rune("A"[0])
	fmt.Println(le1)

	w1 := "Hello"
	le2 := rune(w1[0])
	fmt.Println(le2)

	for i := 65; i <= 122; i++ {
		fmt.Println(i, " - ", string(i), " - ", i%12)
	}
	fmt.Println("hash func ...")
	n := hashBucket("Go", 12)
	fmt.Println(n)
	fmt.Println("Go"[0], int("Go"[0])%12)
	fmt.Println("")
	fmt.Println("---------------")
	res, err := http.Get("https://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	bs, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bs)

	const input = "It is not the critic who counts; not the man who points out how the strong man stumbles."
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "read :", err)
	}
	fmt.Println("-----------------")
	res1, err := http.Get("https://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	scanner1 := bufio.NewScanner(res1.Body)
	defer res1.Body.Close()
	scanner1.Split(bufio.ScanWords)
	for scanner1.Scan() {
		fmt.Println(scanner1.Text())
	}

	fmt.Println("---------------")
	buck := make([]int, 1)
	fmt.Println(buck[0])
	buck[0] = 42
	fmt.Println(buck[0])
	buck[0]++
	fmt.Println(buck[0])

	fmt.Println("---------------")

	res2, err := http.Get("https://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	scanner2 := bufio.NewScanner(res2.Body)
	defer res2.Body.Close()
	scanner2.Split(bufio.ScanWords)
	buck2 := make([]int, 200)
	for scanner2.Scan() {
		n2 := hashBucket2(scanner2.Text())
		buck2[n2]++
	}
	fmt.Println(buck2[65:123])

}

func hashBucket(w string, bus int) int {
	le := int(w[0])
	bu := le % bus
	return bu
}

func hashBucket2(w string) int {
	return int(w[0])
}
