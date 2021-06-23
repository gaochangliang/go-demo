package main

import "fmt"

func main() {
	a := 41
	fmt.Println("a - ", a)
	fmt.Println("a address -", &a)
	var b = &a
	fmt.Println("b- ", b)
}

// the above code makes b a pointer to the memory address where an int is stored
// b is of type "int pointer"
// *int -- the * is part of the type -- b is of type *int
