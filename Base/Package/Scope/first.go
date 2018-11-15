package main

import "fmt"

func main() {
	x, y := get()
	if x == y {
		fmt.Println("equles")
	} else {
		fmt.Println(x, y)
	}
}

func get() (int, int) {
	return 0, 0
}
