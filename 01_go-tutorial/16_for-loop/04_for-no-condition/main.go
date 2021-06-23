package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println("i", i)
		i++
	}
}

// It's a dead end
// The loop will not terminate
