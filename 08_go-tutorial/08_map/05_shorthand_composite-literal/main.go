package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"alice": "hello",
		"bob":   "good morning",
	}
	fmt.Println("myGreeting is", myGreeting)
}
