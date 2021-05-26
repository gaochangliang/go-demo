package main

import "fmt"

func main() {
	var x [58]string

	for i := 64; i < 122; i++ {
		x[i-64] = string(i)
	}

	fmt.Println(x)
	fmt.Println(x[42])
}

//Note that the subscript only goes to 57
