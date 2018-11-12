package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	str := "hello"
	fmt.Fprint(w, str)
}
