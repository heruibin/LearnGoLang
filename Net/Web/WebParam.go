package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/para.do", params)
	http.ListenAndServe("localhost:8080", nil)
}

func params(w http.ResponseWriter, r *http.Request) {

	fmt.Println("HEADER PARAMS:")

	for k, v := range r.Header {
		fmt.Println(k, "=", v)
	}

	fmt.Println("FORM PARAMS:")

	for k, v := range r.Form {
		fmt.Println(k, "=", v)
	}

	fmt.Println("URL PARAMS:", r.URL.RawQuery)
	fmt.Println("URL FORCE QUERY：",r.URL.ForceQuery)

	for k, v := range r.URL.Query() {
		fmt.Println(k, "=", v)
	}

	fmt.Println("END")

	w.Write([]byte("0"))
}
