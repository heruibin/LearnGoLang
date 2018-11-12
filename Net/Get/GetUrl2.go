package main

import (
	"net/http"
	"log"
	"os"
	"io"
	"fmt"
	"strings"
	"bufio"
)

func main() {

	input := bufio.NewScanner(os.Stdin)

	input.Scan()

	url := input.Text()

	if !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("error:", err)
		os.Exit(1)
	}

	fmt.Println("HTTP CODE:", resp.Status, "\t", resp.StatusCode)

	size, err := io.Copy(os.Stdout, resp.Body)

	resp.Body.Close()

	fmt.Println("\r\nSize:", size)

}
