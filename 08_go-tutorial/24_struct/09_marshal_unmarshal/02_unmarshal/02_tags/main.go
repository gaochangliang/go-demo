package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int `json:"wisdom score"`
}

func main() {
	var p1 person

	fmt.Println("p1 first", p1.First)
	fmt.Println("p1 last", p1.Last)
	fmt.Println("p1 age", p1.Age)

	//key Not case-sensitive
	bs := `{"First":"kobe","Last":"bryant","wisdom score":20}`
	json.Unmarshal([]byte(bs), &p1)
	fmt.Println("----------------")
	fmt.Println("p1 first", p1.First)
	fmt.Println("p1 last", p1.Last)
	fmt.Println("p1 age", p1.Age)
	fmt.Printf("%T \n", p1)
}
