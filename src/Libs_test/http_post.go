package main

import (
	"bytes"
	json2 "encoding/json"
	"fmt"
	"net/http"
)

type Person3 struct {
	UserId   string `json:"id"`
	Username string `json:"name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
}

func main() {
	person := Person3{
		UserId:   "120",
		Username: "jack",
		Age:      18,
		Address:  "usa",
	}

	json, _ := json2.Marshal(person)
	reader := bytes.NewReader(json)
	res, err := http.Post("https://golang.org", "application/json;charset=utf-8", reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

}
