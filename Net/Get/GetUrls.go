package main

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"time"
)

func main() {

	start := time.Now()

	urls := []string{"https://www.baidu.com", "https://www.sohu.com", "https://www.linglongtech.com"}

	ch := make(chan string)

	for _, url := range urls {
		go fetch(url, ch)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-ch)
	}

	cost := time.Since(start).Seconds()

	fmt.Println("ALL COST:", cost)
}

func fetch(url string, ch chan<- string) {

	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	cost := time.Since(start).Seconds()

	result := fmt.Sprintf("%s \t %.2f", url, cost)

	ch <- result
}
