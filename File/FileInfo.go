package main

import (
	"os"
	"log"
	"fmt"
)

func main() {
	var fileInfo os.FileInfo
	var error error
	fileInfo, error = os.Stat("D://gofile//test.txt")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
}
