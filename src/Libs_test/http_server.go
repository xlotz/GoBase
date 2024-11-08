package main

import (
	"fmt"
	"net/http"
)

func main() {
	//http.Handle("/index", &MyHandler{})
	//http.ListenAndServe(":8080", nil)

	http.HandleFunc("/index", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(responseWriter, "index")
	})

	http.ListenAndServe(":8080", nil)
}

type MyHandler struct {
}

func (h *MyHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("my implement")

}
