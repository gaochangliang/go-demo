package main

import "fmt"

func main() {
	c := make(chan int)

	/*
	   go func() {
	   		c <- 5
	   	}()
	*/

	c <- 5
	fmt.Println(<-c)
}
