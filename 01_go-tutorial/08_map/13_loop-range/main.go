package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"alice": "hello",
		"bob":   "good morning",
		"trudy": "haha",
	}

	if val, exist := myGreeting["alice"]; exist {
		fmt.Println("book exist", exist)
		fmt.Println("key - alice val -", val)
	} else {
		fmt.Println("book exist", exist)
		fmt.Println("key alice doesn't exist.")
	}
}
