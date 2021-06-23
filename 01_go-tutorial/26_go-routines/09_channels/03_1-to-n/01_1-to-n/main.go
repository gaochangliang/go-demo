package main

import "fmt"

func main() {

	n := 10

	done := make(chan bool)

	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	for i := 0; i < n; i++ {
		go func() {
			for v := range c {
				fmt.Println("value", v)
			}
			done <- true
		}()
	}

	for i := 0; i < n; i++ {
		<-done
	}
}
