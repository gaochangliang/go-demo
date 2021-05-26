package main

import "fmt"

func main() {

	defer hello()
	world()
}

func hello() {
	fmt.Println("hello")
}

func world() {
	fmt.Println("world")
}
