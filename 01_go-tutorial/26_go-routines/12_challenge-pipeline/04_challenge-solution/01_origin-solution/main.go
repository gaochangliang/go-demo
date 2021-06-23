package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	out := gen()
	f := factorial(out)
	for v := range f {
		fmt.Println("f", v)
	}
}

func gen() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for k := 3; k < 13; k++ {
				out <- k
			}
		}
		close(out)
	}()

	return out
}

func factorial(c chan int) chan int {
	out := make(chan int)
	go func() {
		for v := range c {
			total := fact(v)
			out <- total
		}

	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := 2; i <= n; i++ {
		total *= i
	}
	return total
}
