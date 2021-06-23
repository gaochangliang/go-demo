package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"alice": "hello",
		"bob":   "good morning",
	}
	fmt.Println("len myGreeting is", len(myGreeting))

	myGreeting["trudy"] = "haha"

	fmt.Println("len after add myGreeting  is", len(myGreeting))
}
