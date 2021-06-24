package main

import "fmt"

type myType []int

func main() {
	var x myType = []int{32, 33}
	fmt.Println("x", x)
	fmt.Printf("%T \n", x)
	fmt.Printf("%T", []int(x))

}
