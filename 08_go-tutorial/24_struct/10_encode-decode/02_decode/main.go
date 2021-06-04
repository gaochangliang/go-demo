package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var p1 person
	rdr := strings.NewReader(`{"First":"kobe", "Last":"bryant", "Age":41}`)
	json.NewDecoder(rdr).Decode(&p1)

	fmt.Println("p1 first", p1.First)
	fmt.Println("p1 last", p1.Last)
	fmt.Println("p1 first", p1.Age)

}
