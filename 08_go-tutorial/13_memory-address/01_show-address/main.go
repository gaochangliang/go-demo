package main

import "fmt"

func main() {
	name := "kobe"
	fmt.Println("name - ", name)
	fmt.Println("name memory address - ", &name)
	fmt.Printf("%d \n", &name)
}
