package main

import "fmt"

func main() {
	n := gen()
	out := factorial(n)
	for v := range out {
		fmt.Println("total", v)
	}

}

func gen() chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := 1; i <= n; i++ {
		total *= i
	}
	return total
}

func factorial(n chan int) chan int {
	out := make(chan int)
	go func() {
		for v := range n {
			out <- fact(v)
		}
		close(out)
	}()
	return out
}
