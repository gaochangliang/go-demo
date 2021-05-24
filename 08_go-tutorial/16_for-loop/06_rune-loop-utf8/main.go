package main

import "fmt"

func main() {
	for i := 100; i <= 250; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
}
