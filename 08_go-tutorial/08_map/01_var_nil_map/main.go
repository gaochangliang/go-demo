package main

import "fmt"

func main() {
	var myGreeting map[string]string
	fmt.Println(myGreeting)
	fmt.Println("myGreeting = nil", myGreeting == nil)

	//myGreeting["alice"] = "hello"
	//myGreeting["bob"] = "good morning"

	//add these lines
	//you will get  panic: assignment to entry in nil map
}
