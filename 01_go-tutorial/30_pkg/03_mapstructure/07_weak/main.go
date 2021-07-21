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
		"age":  "18",
		"job":  "basket",
	}

	var p Person
	err := mapstructure.WeakDecode(m, &p)

	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Printf("%+v", p) //%+v可以打印结构体的key value

}
