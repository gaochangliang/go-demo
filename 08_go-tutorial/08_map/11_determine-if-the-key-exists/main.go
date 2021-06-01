package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"alice": "hello",
		"bob":   "good morning",
		"trudy": "haha",
	}

	fmt.Println("greeting", myGreeting)

	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}
}
