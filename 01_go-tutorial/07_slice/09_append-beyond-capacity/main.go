package main

import "fmt"

func main() {
	greeting := make([]string, 3, 5)

	greeting[0] = "good morning"
	greeting[1] = "Bonjour!"
	greeting[2] = "suprabadham"

	fmt.Println("len - ", len(greeting), "cap", cap(greeting))

	greeting = append(greeting, "buenos dias!")
	fmt.Println("len - ", len(greeting), "cap", cap(greeting))

	greeting = append(greeting, "Suprabadham")
	fmt.Println("len - ", len(greeting), "cap", cap(greeting))

	greeting = append(greeting, "Zǎo'ān")
	fmt.Println("len - ", len(greeting), "cap", cap(greeting))

	greeting = append(greeting, "Ohayou gozaimasu")
	fmt.Println("len - ", len(greeting), "cap", cap(greeting))

	greeting = append(greeting, "gidday")
	fmt.Println("len - ", len(greeting), "cap", cap(greeting))

	fmt.Println("greeting[3]", greeting[3])
	fmt.Println("greeting[4]", greeting[4])
	fmt.Println("greeting[5]", greeting[5])
	fmt.Println("greeting[6]", greeting[6])

}
