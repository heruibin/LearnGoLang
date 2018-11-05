package main

import (
	"io/ioutil"
	"fmt"
)

func main() {

	data, err := ioutil.ReadFile("D://gofile//test.txt")
	if err != nil {
		fmt.Println("Error!")
	}

	lines := string(data)

	fmt.Println(lines)


}
