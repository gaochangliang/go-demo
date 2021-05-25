package main

import (
	"fmt"
)

func main() {
	mySlice := make([]int, 0, 3)

	fmt.Println("-----")
	fmt.Println("mySlice is ", len(mySlice))
	fmt.Println("length is ", len(mySlice))
	fmt.Println("cap is", cap(mySlice))
	fmt.Println("-----")

	for i := 0; i <= 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("len - ", len(mySlice), "cap - ", cap(mySlice), "value - ", mySlice[i])
	}
}

//cap 2X expansion
