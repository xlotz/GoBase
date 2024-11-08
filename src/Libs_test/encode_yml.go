package main

/**

yml
*/
//import "github.com/go-yaml/yaml"

import (
	"fmt"
	yaml "gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Database string `yaml:"database"`
	Url      string `yaml:"url"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func main() {

	// 序列化
	fmt.Println("序列化...")
	config := Config{
		Database: "oracle",
		Url:      "localhost",
		Port:     3306,
		Username: "root",
		Password: "123456",
	}

	out, err := yaml.Marshal(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))

	fmt.Println("反序列化...")
	bytes, err := os.ReadFile("./src/file/config.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	var config2 Config
	err2 := yaml.Unmarshal(bytes, &config2)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(config2)

}
