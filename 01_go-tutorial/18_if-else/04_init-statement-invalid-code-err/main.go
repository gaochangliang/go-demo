package main

import "fmt"

func main() {
	b := true
	if name := "kobe"; b {
		fmt.Println("name %s", name)
	}

	fmt.Println("name2", name)
}

//The if statement is a block of code, locally scoped and not accessible externally
