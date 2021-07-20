package main

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"log"
)

type Person struct {
	//Using the tag mapstructure: "username" maps the Person's Name field to username,
	//and we need to set username in the JSON string in order to parse it correctly
	Name string `mapstructure:"username"`
	Age  int
	Job  string
}

type Cat struct {
	Name  string
	Age   int
	Breed string
}

func main() {
	//We capitalized the Username and the Age in the first JSON string,
	//but it did not affect the decoding result. mapstructure handles field mapping case-insensitively
	datas := []string{`
    { 
      "type": "person",
      "Username":"kobe",
      "Age":41,
      "job": "basketball"
    }
  `,

		`
    { 
      "type": "person",
      "name":"jordan",
      "age":23,
      "job": "basketball"
    }
  `,

		`
    {
      "type": "cat",
      "username": "kitty",
      "age": 1,
      "breed": "Ragdoll"
    }
  `,
	}

	for _, data := range datas {
		var m map[string]interface{}

		//use json.Unmarshal to decode the byte stream to the map[string]interface{} type.
		//map[age:41 job:basketball name:kobe type:person]
		err := json.Unmarshal([]byte(data), &m)
		if err != nil {
			log.Fatal("json.Unmarshal", err)
		}

		//Type Assertion
		//read the type field inside. Based on the value of the type field,
		//the JSON string is then decoded into Person and Cat values using mapstructure.Decode, and output
		switch m["type"].(string) {
		case "person":
			var p Person
			mapstructure.Decode(m, &p)
			log.Println("person", p)

		case "cat":
			var cat Cat
			mapstructure.Decode(m, &cat)
			log.Println("cat", cat)
		}
	}

}

//
