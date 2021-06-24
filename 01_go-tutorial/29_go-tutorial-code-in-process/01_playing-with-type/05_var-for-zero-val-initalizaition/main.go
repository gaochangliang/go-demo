package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var p person
	fmt.Println("p.name", p.name)
	fmt.Println("p.age", p.age)
	fmt.Printf("%T\n", p)
}

// always use var to create and
// initialize a variable to its zero value
