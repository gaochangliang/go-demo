package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println("mySlice len", len(mySlice), "cap", cap(mySlice))

	myAnotherSlice := []int{6, 7, 8, 9}

	mySlice = append(mySlice, myAnotherSlice...)
	fmt.Println("mySlice update len", len(mySlice), "cap", cap(mySlice))
	fmt.Println(mySlice)

}
