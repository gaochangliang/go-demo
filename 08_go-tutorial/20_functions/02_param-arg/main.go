package main

import "fmt"

func main() {
	greet("kobe")
	greet("CR")
}

func greet(name string) {
	fmt.Println(" name is ", name)
}

// greet is declared with a parameter
// when calling greet, pass in an argument
