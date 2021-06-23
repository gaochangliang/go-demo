package main

import (
	"encoding/json"
	"os"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := person{
		First:       "kobe",
		Last:        "Bryant",
		Age:         41,
		notExported: 1,
	}

	json.NewEncoder(os.Stdout).Encode(p1)

}
