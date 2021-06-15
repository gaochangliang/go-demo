package main

import "fmt"

func main() {
	done := make(chan bool)
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		for v := range c {
			fmt.Println("value", v)
		}
		done <- true
	}()

	go func() {
		for v := range c {
			fmt.Println("value", v)
		}
		done <- true
	}()

	<-done
	<-done

}
