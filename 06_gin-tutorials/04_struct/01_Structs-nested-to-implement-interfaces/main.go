package main

import "fmt"

func main() {

	var p person
	Print(&p)

	var p2 person2
	Print(&p2)

	var p3 person3
	Print(&p3.p)

}

type person struct {
}

type person2 struct {
	person
}

type person3 struct {
	p person
}

func (p *person) testP() string {
	return "hello"
}

func Print(i in1) {
	fmt.Println(1)
}

type in1 interface {
	testP() string
}
