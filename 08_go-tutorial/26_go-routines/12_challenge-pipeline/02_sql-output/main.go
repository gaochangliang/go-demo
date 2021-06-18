package main

import "fmt"

func main() {
	for v := range sq(sq(gen(2, 3))) {
		fmt.Println(v)
	}
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
