package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name string
	Age  int
	Job  string `mapstructure:",omitempty"`
}

func main() {
	var p = Person{
		Name: "kobe",
		Age:  18,
	}

	var m map[string]interface{}
	mapstructure.Decode(p, &m)

	data, _ := json.Marshal(m)
	fmt.Println("data", string(data))

	//If we don't want to display fields without data,
	//we can add mapstructure:",omitempty" to the corresponding field of the structure

}
