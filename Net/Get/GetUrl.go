package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"os"
)

func main() {
	url := "https://www.baidu.com/"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("error:", err)
		os.Exit(1)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("error:", err)
		os.Exit(1)
	}
	respStr := string(respBytes)

	resp.Body.Close()

	fmt.Println(respStr)
}
