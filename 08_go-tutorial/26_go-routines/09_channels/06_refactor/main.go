package main

import "fmt"

func main() {
	c := incrementor()
	for k := range puller(c) {
		fmt.Println(k)
	}

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
