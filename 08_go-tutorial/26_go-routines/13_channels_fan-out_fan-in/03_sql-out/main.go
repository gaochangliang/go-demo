package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3)
	c1 := sq(in)
	c2 := sq(in)
	for v := range merge(c1, c2) {
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
	fmt.Println("sq c type %T", c)
	out := make(chan int)
	go func() {
		for n := range c {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...chan int) chan int {
	fmt.Println("merge cs type %T", cs)
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(c chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(c)
	}
	//Firstly, this is a pipeline, secondly, make sure that all the chan loops are executed before closing
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
