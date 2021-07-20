package main

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"log"
)

type Person struct {
	Name string
	Age  int
	Job  string
}

type Cat struct {
	Name  string
	Age   int
	Breed string
}

func main() {
	datas := []string{`
    { 
      "type": "person",
      "name":"kobe",
      "age":41,
      "job": "basketball"
    }
  `,
		`
    {
      "type": "cat",
      "name": "kitty",
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

/*

We have defined two structures Person and Cat, which have slightly different fields.
Now, the JSON string we agreed to communicate with has a type field in it.
When the value of type is person, the JSON string represents data of type Person.
When the value of type is cat, the JSON string represents data of type Cat.

*/

//In fact, Google Protobuf usually uses this approach as well.
//Add a message ID or fully qualified message name to the protocol.
//When the receiver receives the data, it first reads the protocol ID or fully qualified message name.
//Then Protobuf's decode method is called to decode it into the corresponding Message structure.
//From this point of view, mapstructure can also be used for network message decoding
//if you don't consider performance
