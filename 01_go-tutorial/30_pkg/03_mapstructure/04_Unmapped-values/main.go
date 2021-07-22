package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

//If there is an unmapped value in the source data (i.e. no corresponding field in the structure),
//mapstructure will ignore it by default.
//We can define a field in the structure and set the mapstructure:",remain" tag for it.
//This way unmapped values will be added to this field.
//Note that this field can only be of type map[string]interface{} or map[interface{}]interface{}
type Person struct {
	Name  string
	Age   int
	Job   string
	Other map[string]interface{} `mapstructure:",remain"`
}

func main() {
	data := `
		{
		  "name": "kobe",
		  "age":18,
		  "job":"basketball",
		  "height":"1.8m",
		  "handsome": true
		}
	`

	var m map[string]interface{}
	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		fmt.Println("err", err)
	}

	var p Person
	mapstructure.Decode(m, &p)
	fmt.Println("p", p)
	fmt.Println("other", p.Other)

}
