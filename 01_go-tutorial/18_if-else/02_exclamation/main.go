package main

import "fmt"

func main() {
	if !true {
		fmt.Println("this did not ran")
	}

	if !false {
		fmt.Println("this ran")
	}
}
