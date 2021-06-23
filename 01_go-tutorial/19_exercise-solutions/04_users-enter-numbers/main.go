package main

import "fmt"

func main() {
	var numOne, numTwo int
	fmt.Println("please enter a large number")
	fmt.Scan(&numOne)
	fmt.Println("please enter a large number")
	fmt.Scan(&numTwo)

	fmt.Println(numOne, "%", numTwo, " = ", numOne%numTwo)
}

//use numOne not num1
