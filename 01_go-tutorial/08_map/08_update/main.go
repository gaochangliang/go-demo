package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"alice": "hello",
		"bob":   "good morning",
	}
	fmt.Println("len myGreeting is", myGreeting)

	myGreeting["alice"] = "haha"

	fmt.Println("update myGreeting", myGreeting)
}
