package main

import "fmt"

func main() {
	var i int
	p := &i
	fmt.Println(p)
	*p = 3
	fmt.Println(i)


}
