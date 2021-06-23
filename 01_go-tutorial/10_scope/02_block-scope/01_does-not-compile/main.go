package main

import "fmt"

func main() {
	var x = 42
	fmt.Println("main", x)
	foo()
}

// no access to x
// this does not compile
func foo() {
	fmt.Println("foo", x)
}
