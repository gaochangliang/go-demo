package main

import "fmt"

//联系人
type contact struct {
	greeting string
	name     string
}

func main() {

}

func switchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}
