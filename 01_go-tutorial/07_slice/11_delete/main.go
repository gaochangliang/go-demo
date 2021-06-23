package main

import "fmt"

func main() {
	mySlice := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

	//delete wednesday
	mySlice = append(mySlice[:2], mySlice[3:]...)
	fmt.Println(mySlice)
}
