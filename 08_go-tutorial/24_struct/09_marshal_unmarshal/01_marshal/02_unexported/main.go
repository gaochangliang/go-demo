package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	first       string
	last        string
	age         int
	notExported int //Lowercase beginnings cannot be exported
}

func main() {
	p1 := person{
		first:       "kobe",
		last:        "Bryant",
		age:         41,
		notExported: 1,
	}

	bs, _ := json.Marshal(p1)
	fmt.Println("bs", bs)
	fmt.Println("bs string", string(bs))
	fmt.Printf("%T \n", bs)

}
