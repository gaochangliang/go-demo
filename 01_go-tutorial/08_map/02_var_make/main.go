package main

import "fmt"

func main() {
	var myGreeting = make(map[string]string)
	myGreeting["alice"] = "hello"
	myGreeting["bob"] = "good morning"
	fmt.Println("myGreeting is", myGreeting)
}
