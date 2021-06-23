package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p := person{
		name: "kobe",
		age:  41,
	}
	fmt.Printf("%T\n", p)
	fmt.Println(p.name)
	fmt.Println(p.age)
	fmt.Println(p)
}
