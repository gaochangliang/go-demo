package main

import "fmt"

func main() {
	n := factorial(4)
	for v := range n {
		fmt.Println("total", v)
	}

}

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := 1; i <= n; i++ {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}

/*
CHALLENGE #1:
-- Use goroutines and channels to calculate factorial

CHALLENGE #2:
-- Why might you want to use goroutines and channels to calculate factorial?
---- Formulate your own answer, then post that answer to this discussion area: https://goo.gl/flGsyX
---- Read a few of the other answers at the discussion area to see the reasons of others
*/
