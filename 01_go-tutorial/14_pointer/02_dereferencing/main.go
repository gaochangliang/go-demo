package main

import "fmt"

func main() {
	a := 43

	fmt.Println("a value is ", a)
	fmt.Println("a address is ", &a)

	var b = &a
	fmt.Println("b value is ", b)
	fmt.Println("&b is ", &b)
}
