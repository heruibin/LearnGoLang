package main

import (
	"os"
	"log"
)

func main() {
	var newFile *os.File
	var error error
	newFile, error = os.Create("D://gofile//test.txt")
	if error != nil {
		log.Fatal("error:", error)
	}
	log.Println("success")
	newFile.Close()
}
