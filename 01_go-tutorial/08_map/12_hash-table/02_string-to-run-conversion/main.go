package main

import "fmt"

func main() {
	letter := rune("A"[0])
	fmt.Printf("%T - %v", letter, letter)
}
