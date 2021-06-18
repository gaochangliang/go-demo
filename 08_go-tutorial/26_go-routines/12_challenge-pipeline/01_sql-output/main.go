package main

import "fmt"

func main() {
	c := gen(2, 3)
	total := sq(c)
	fmt.Println("total 1", <-total)
	fmt.Println("total 2", <-total)
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()
	return out
}

func sq(c chan int) chan int {
	total := make(chan int)
	go func() {
		for v := range c {
			value := v * v
			total <- value
		}
		close(total)
	}()
	return total
}
