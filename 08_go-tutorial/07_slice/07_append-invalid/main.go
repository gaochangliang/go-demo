package main

import "fmt"

func main() {
	greeting := make([]string, 3, 5)

	greeting[0] = "good morning"
	greeting[1] = "Bonjour!"
	greeting[2] = "suprabadham"
	greeting[3] = "buenos dias!" //error

	fmt.Println(greeting[2])
}

// The length must be specified when slice is made
// Subscripts must be less than the length
