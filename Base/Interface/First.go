package main

import "fmt"

type Person struct{
	name string
}

type Dog struct{
	name string
}

type Animal interface{
	say()
}

func (d Dog) say(){
	fmt.Println(d.name,":i am dog,wang wang")
}

func (p Person) say(){
	fmt.Println(p.name,":i am human,yapee")
}

func (p Person) say2(){
	fmt.Println(p.name,":i am human,yapee")
}

func (p *Person) say3(){
	fmt.Println(p.name,":i am human,yapee")
}

func main() {
	var pb Animal
	pa := Dog{"snoopy"}
	pb = Person{"nana"}
	pa.say()
	pb.say()
	pc := Person{"yoyo"}
	pc.say2()
	pc.say3()
}
