package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"alice": "hello",
		"bob":   "good morning",
	}

	myGreeting["trudy"] = "haha"
	fmt.Println("myGreeting is", myGreeting)
}
