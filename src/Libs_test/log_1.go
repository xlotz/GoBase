package main

import "log"

func main() {

	log.Println("-------------------")
	log.SetPrefix("[log] ")
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Lmsgprefix | log.Ldate | log.Ltime)
	log.Println("日志")
	log.Panicln("panic日志")
	log.Fatalln("错误日志")
}
