package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"alice": "hello",
		"bob":   "good morning",
		"trudy": "haha",
	}

	fmt.Println("myGreeting one", myGreeting)
	delete(myGreeting, "kobe")
	fmt.Println("myGreeting two", myGreeting)
}
