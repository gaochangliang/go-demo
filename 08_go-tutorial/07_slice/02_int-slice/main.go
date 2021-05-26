package main

import "fmt"

func main() {
	var mySlice = []int{1, 2, 3, 4, 5}

	for i, v := range mySlice {
		fmt.Println(i, " - ", v)
	}
}
