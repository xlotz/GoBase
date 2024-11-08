package main

import (
	"encoding/json"
	"fmt"
)

/**
json
*/

type Person2 struct {
	UserId   string `json:"id"`
	Username string `json:"name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
}

func main() {

	fmt.Println("序列化...")
	person := Person2{
		UserId:   "120",
		Username: "jack",
		Age:      18,
		Address:  "usa",
	}

	bytes, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))

	fmt.Println("序列化，缩进...")
	bytes2, err2 := json.MarshalIndent(person, "", "\t")
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(string(bytes2))

	fmt.Println("反序列化...")

	person2 := Person2{}
	jsonStr := "{\"id\":\"120\",\"name\":\"jack\",\"age\":18,\"address\":\"usa\"}\n"
	err3 := json.Unmarshal([]byte(jsonStr), &person2)
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	fmt.Printf("%+v", person2)
}
