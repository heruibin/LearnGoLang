package main

import (
	"os"
	"log"
)

func main() {
	error := os.Truncate("D://gofile/test.txt", 10)
	if error != nil {
		log.Fatal(error)
	}
}
