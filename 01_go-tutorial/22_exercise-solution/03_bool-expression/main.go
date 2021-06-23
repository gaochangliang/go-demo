package main

import "fmt"

func main() {
	largest := max(100, 233, 344, 500, 230, 310)
	fmt.Println("the largest number is", largest)
}

func max(numbers ...int) int {
	var largest int
	for _, value := range numbers {
		if value > largest {
			largest = value
		}
	}
	return largest
}

//Find all parameters that pass the maximum
