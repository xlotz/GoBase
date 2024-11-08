package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type person1 struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {

	p1 := person1{"A", "B", 20, 007}
	json.NewEncoder(os.Stdout).Encode(p1)

	var p2 person1
	rdr := strings.NewReader(`{"First":"AA", "Last":"BB", "Age":20}`)
	json.NewDecoder(rdr).Decode(&p2)
	fmt.Println(p2.First, p2.Last, p2.Age)
}
