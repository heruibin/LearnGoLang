package main

import (
	"os"
	"log"
)

func main() {
	var error error
	error = os.Rename("D://gofile//test.txt", "D://gofile//test2.txt")
	if error != nil {
		log.Fatal(error)
	}

}
