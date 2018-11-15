package main

import "fmt"

type Abc float64
type Def float64

func main() {
	var abc Abc
	var def Def
	var aaa float64
	fmt.Println(abc == 0)
	fmt.Println(abc == Abc(aaa))
	fmt.Println(abc == Abc(def))

}
