package main

import "fmt"

type person struct {
	first  string
	last   string
	number int
}

func main() {
	p1 := person{
		first:  "jordan",
		last:   "Michael",
		number: 23,
	}
	p2 := person{
		first:  "kobe",
		last:   "bryant",
		number: 24,
	}
	fmt.Println("p1", p1.first, p1.last, p1.number)
	fmt.Println("p1", p2.first, p2.last, p2.number)
}
