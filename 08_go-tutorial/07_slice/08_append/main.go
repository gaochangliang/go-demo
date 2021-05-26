package main

import "fmt"

func main() {
	greeting := make([]string, 3, 5)

	greeting[0] = "good morning"
	greeting[1] = "Bonjour!"
	greeting[2] = "suprabadham"

	greeting = append(greeting, "buenos dias!")

	fmt.Println(greeting[3])
}
