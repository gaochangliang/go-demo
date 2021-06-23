package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First       string `json:"first"`
	Last        string `json:"-"` //  - Indicates that this field is not exported
	Age         int    `json:"age"`
	notExported int    `json:"not_exported"`
}

func main() {
	p1 := person{
		First:       "kobe",
		Last:        "Bryant",
		Age:         41,
		notExported: 1,
	}

	bs, _ := json.Marshal(p1)
	fmt.Println("bs", bs)
	fmt.Println("bs string", string(bs))
	fmt.Printf("%T \n", bs)

}
