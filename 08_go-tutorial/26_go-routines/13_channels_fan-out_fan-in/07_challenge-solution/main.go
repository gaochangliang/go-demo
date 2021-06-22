package main

import (
	"fmt"
	"sync"
)

var wordId int
var publishId int

func main() {
	in := gen()

	c1 := fact(in)
	c2 := fact(in)
	c3 := fact(in)
	c4 := fact(in)
	c5 := fact(in)
	c6 := fact(in)
	c7 := fact(in)
	c8 := fact(in)
	c9 := fact(in)
	out := merge(c1, c2, c3, c4, c5, c6, c7, c8, c9)
	for v := range out {
		fmt.Println("range", v)
	}

}

func gen() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func fact(c chan int) chan int {
	out := make(chan int)
	go func() {
		for v := range c {
			out <- factorial(v)
		}
		close(out)
	}()
	return out
}

func factorial(n int) int {
	total := 1
	for i := 2; i <= n; i++ {
		total *= i
	}
	return total
}

func merge(c ...chan int) chan int {
	fmt.Println("merge %T", c)
	var wg sync.WaitGroup
	out := make(chan int)
	fact := func(c chan int) {
		for v := range c {
			out <- v
		}
		wg.Done()
	}
	//The add here must be placed outside
	wg.Add(len(c))

	for _, v := range c {
		go fact(v)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
