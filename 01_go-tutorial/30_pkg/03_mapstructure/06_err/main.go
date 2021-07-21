package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name string
	Age  int
	Job  string
}

func main() {
	m := map[string]interface{}{
		"name": 18,
		"age":  "kobe",
		"job":  "basket",
	}

	var p Person
	err := mapstructure.Decode(m, &p)

	if err != nil {
		fmt.Println("err", err)
	}

}
