package main

import (
	"fmt"
	"log"
	http2 "net/http"
)

func main() {

	http2.HandleFunc("/", indexHandler)
	http2.HandleFunc("/hello", helloHandler)
	log.Fatal(http2.ListenAndServe(":8888", nil))
}

func indexHandler(w http2.ResponseWriter, req *http2.Request) {
	fmt.Fprintf(w, "URL.Path=%q\n", req.URL.Path)
}

func helloHandler(w http2.ResponseWriter, req *http2.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
