package main

import (
	"fmt"
)

func main() {
	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	for i, v := range greeting {
		fmt.Println("i--v", i, v)
	}

	for i := 0; i < len(greeting); i++ {
		fmt.Println("v", greeting[i])
	}
}

//cap 2X expansion
