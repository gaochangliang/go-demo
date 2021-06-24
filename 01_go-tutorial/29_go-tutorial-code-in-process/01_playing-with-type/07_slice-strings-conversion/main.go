package main

import "fmt"

type mySentence []string

func main() {
	var myS mySentence = []string{"hello", "world"}
	fmt.Println("myS", myS)
	fmt.Printf("mySentence   %T\n", myS)
	fmt.Printf("mySentence []string   %T\n", []string(myS))
}
