package main

import "fmt"

type foo int

func main() {
	var kobeAge foo
	kobeAge = 41
	fmt.Printf("%T - %v\n", kobeAge, kobeAge)

	var myAge int
	myAge = 100
	fmt.Printf("%T %v\n", myAge, myAge)

	// this doesn't work
	//fmt.Println(kobeAge + myAge)

	fmt.Println("sum", int(kobeAge)+myAge)

}
