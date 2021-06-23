package main

import "fmt"

type mySentence string

func main() {
	var myS mySentence = "hello world"
	fmt.Println("string", myS)
	fmt.Printf("%T\n", myS)
}
