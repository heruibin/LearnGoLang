package main

import "fmt"

type Addr func(a int, b int) int

func add(a int, b int) int {
	return a + b
}

func add2(a int, b int) {
	fmt.Println("add2:", a+b)
}

func main() {
	var addr Addr
	addr = add
	fmt.Println(addr(1,2))
}
