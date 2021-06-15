package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)
	n := 10

	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			done <- true
		}()
	}

	go func() {
		//The for loop here ensures that all 10 goroutines above are executed
		//and that the closed channel can read the data
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	//channel  here must close Otherwise panic
	for n := range c {
		fmt.Println(n)
	}
}
