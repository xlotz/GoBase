package main

import "net/http"

func main() {

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://golang.org", nil)
	req.Header.Add("Authorization", "123456")
	resq, _ := client.Do(req)
	defer resq.Body.Close()
}
