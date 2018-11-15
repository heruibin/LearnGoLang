package main

import "fmt"

const (
	GB int = 1 << (iota * 10)
	MB int = 1 << (iota * 10)
	KB int = 1 << (iota * 10)
)

const (
	a int = iota
	b
	_
	c
	d string = "abc"
)

const(
	aa int = iota * 10
	bb int = iota * 10
	cc
	dd
)

func main() {
	fmt.Println(GB, MB, KB)
	fmt.Println(a, b, c, d)
	fmt.Println(aa, bb, cc, dd)
	fmt.Println(1 << 20)

	fmt.Println( 3 << 1)
}
