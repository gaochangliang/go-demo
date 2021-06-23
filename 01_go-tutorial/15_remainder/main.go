package main

import "fmt"

func main() {
	x := 9 % 4
	fmt.Println("x value is ", x)
	if x == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}
}
