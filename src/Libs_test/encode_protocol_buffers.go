package main

import (
	p "/Users/xiaoqiang/project/goproject/GoBase/src/proto"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {
	person := p.Person{
		Name:   "wyh",
		Age:    12,
		Gender: p.Gender_FE_MAIL,
	}
	// 序列化
	data, err := proto.Marshal(&person)
	if err != nil {
		fmt.Println(err)
		return
	}

	temp := &p.Person{}
	fmt.Println("proto buffer len: ", len(data), "bytes: ", data)
	err = proto.Unmarshal(data, temp) // 反序列化

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(temp)
}
