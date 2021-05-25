package main

import "fmt"

func main() {
	var result []int
	fmt.Println("result", result)
	fmt.Println("len", len(result), "cap", cap(result))

	mySlice := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println("mySlice is", mySlice)
	fmt.Println("mySlice[2:4]", mySlice[2:4])
	fmt.Println("mySlice[2]", mySlice[2])
	fmt.Println("myString"[2])
}

//Cutting follows the left-close, right-open principle
