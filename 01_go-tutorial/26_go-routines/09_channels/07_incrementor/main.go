package main

import "fmt"

func main() {
	c1 := incrementor()
	c2 := incrementor()
	p1 := puller(c1)
	p2 := puller(c2)
	fmt.Println("p1+p2=", <-p1+<-p2)
}

func incrementor() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		//Note the close here
		close(c)
	}()
	return c
}

func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int
		for v := range c {
			sum += v
		}
		out <- sum
		close(out)
	}()
	return out
}
