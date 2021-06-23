package main

import "fmt"

func main() {
	if false {
		fmt.Println("first print statement")
	} else if true {
		fmt.Println("second print statement")
	} else if false {
		fmt.Println("third print statement")
	} else {
		fmt.Println("fourth print statement")
	}
}

//这里的true false实际使用中都是判断条件
