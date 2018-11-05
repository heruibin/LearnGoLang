package main

import (
	"Algorithm/Sort/BubbleSort"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	arrays := []int{2, 3, 1, 5, 6, 4, 9}
	arrays = bubble.BubbleSort(arrays)
	fmt.Printf("[]int:", arrays)
	fmt.Println()
	fmt.Printf("array:", arrays)
	fmt.Println()
	fmt.Printf("array,slice:", arrays)
	fmt.Println()
}
