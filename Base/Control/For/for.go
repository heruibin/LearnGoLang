package main

import (
	"fmt"
	"strings"
)

func main() {
	var arrays []string = []string{"1", "2", "3", "4"}
	for i := 0; i < len(arrays); i++ {
		fmt.Print(arrays[i], " ")
	}
	fmt.Println()
	for _, i := range arrays {
		fmt.Print(i, " ")
	}
	fmt.Println()
	count := 0
	for {
		fmt.Print(arrays[count], " ")
		count = count + 1
		if count == len(arrays) {
			break
		}
	}
	strings.Join()
}
