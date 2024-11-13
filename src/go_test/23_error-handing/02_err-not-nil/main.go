package main

import (
	"os"
)

func main() {

	/**
	fmt.Println / log.Println /log.Fataln
	*/

	_, err := os.Open("files/no-file.txt")
	if err != nil {
		//fmt.Println("err happened", err)
		//log.Println("err happened", err)
		//log.Fatalln(err)
		panic(err)
	}

}
