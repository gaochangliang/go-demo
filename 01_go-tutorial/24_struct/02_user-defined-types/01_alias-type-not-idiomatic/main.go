package main

import "fmt"

type foo int

func main() {
	var kobeAge foo
	kobeAge = 41
	fmt.Printf("%T - %v\n", kobeAge, kobeAge)
}
