package main

import (
	"encoding/xml"
	"fmt"
)

/**
xml
*/

type Person struct {
	UserId   string `xml:"id"`
	Username string `xml:"name"`
	Age      int    `xml:"age"`
	Address  string `xml:"address"`
}

// func Marshal(v any) ([]byte, error) //xml 序列化
// func MarshalIndent(v any, prefix, indent string) ([]byte, error) //格式化
// func Unmarshal(data []byte, v any) error // 反序列化
func main() {
	// 序列化
	fmt.Println("序列化...")
	person := Person{
		UserId:   "120",
		Username: "jack",
		Age:      18,
		Address:  "usa",
	}

	bytes, err := xml.MarshalIndent(person, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))

	fmt.Println("反序列化...")
	var person2 = Person{
		UserId:   "",
		Username: "",
		Age:      0,
		Address:  "",
	}

	xmlStr := "<Person>			\n		<id>120</id>		\n		<name>tom</name>		\n		<age>20</age>		\n		<address>uuu</address>\n</Person>"
	err2 := xml.Unmarshal([]byte(xmlStr), &person2)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(xmlStr)

}
