package main

import (
	"fmt"
	"gee"
)

func main() {
	person := gee.H{
		"name": "kobe",
		"age":  41,
	}
	fmt.Println(person)
	if val, ok := person["name"]; ok {
		fmt.Printf("name->%s", val)
	}
}
