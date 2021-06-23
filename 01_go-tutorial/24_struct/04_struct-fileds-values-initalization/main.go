package main

import "fmt"

type person struct {
	first  string
	last   string
	number int
}

//notes + here
func (p person) fullName() string {
	return p.first + p.last
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
	fmt.Println("p1 fullName", p1.fullName())
	fmt.Println("p2 fullName ", p2.fullName())
}
