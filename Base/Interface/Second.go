package main

import "fmt"

type A interface {
	aFirst()
	aSecond()
}

type B interface {
	bFirst()
	bSecond()
}

type C interface {
	aFirst()
	aSecond()
	bFirst()
	bSecond()
}

type Impl struct {
	name string
	age  int
}

func (impl Impl) aFirst() {
	fmt.Println("aFirst ",impl.name)
	impl.name = "hualalal"
}

func (impl Impl) aSecond() {
	fmt.Println("aSecond")
}

func (impl Impl) bFirst() {
	fmt.Println("bFirst")
}

func (impl Impl) bSecond() {
	fmt.Println("bSecond")
}

func main() {
	var test C
	test = Impl{"rbhe",28}
	test.aFirst()
	test.bFirst()
}
