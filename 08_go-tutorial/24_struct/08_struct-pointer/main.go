package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := &person{
		name: "kobe",
		age:  41,
	}
	fmt.Println("p1", p1)
	fmt.Printf("%T\n", p1)
	fmt.Println("name - ", p1.name)
	fmt.Println("age - ", p1.age)
}
