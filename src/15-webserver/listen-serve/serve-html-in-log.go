package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello")
}

// from Todd McLeod: https://github.com/GoesToEleven/golang-web-dev
