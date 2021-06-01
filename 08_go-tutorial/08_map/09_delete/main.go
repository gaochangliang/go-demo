package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"alice": "hello",
		"bob":   "good morning",
		"trudy": "haha",
	}
	fmt.Println("myGreeting is", myGreeting)

	delete(myGreeting, "trudy")

	fmt.Println("myGreeting delete trudy ", myGreeting)
}
