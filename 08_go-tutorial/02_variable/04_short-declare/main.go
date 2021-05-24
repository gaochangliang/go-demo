package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "kobe"
	a, b, c := 1, false, "kurry"
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(name, a, b, c)
}
