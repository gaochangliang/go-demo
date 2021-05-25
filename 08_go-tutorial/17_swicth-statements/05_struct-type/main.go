package main

import "fmt"

//联系人
type contact struct {
	greeting string
	name     string
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

func main() {
	switchOnType(7)
	switchOnType("kobe")
	var t = contact{
		greeting: "kobe",
		name:     "kobe",
	}
	switchOnType(t)
	switchOnType(t.greeting)
	switchOnType(t.name)
}
