package main

import "fmt"

func main() {
	n := factorial(4)
	total := 1
	for v := range n {
		total *= v
	}
	fmt.Println("total", total)
}

func factorial(n int) chan int {
	c := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

/*
CHALLENGE #1:
-- Use goroutines and channels to calculate factorial

CHALLENGE #2:
-- Why might you want to use goroutines and channels to calculate factorial?
---- Formulate your own answer, then post that answer to this discussion area: https://goo.gl/flGsyX
---- Read a few of the other answers at the discussion area to see the reasons of others
*/
