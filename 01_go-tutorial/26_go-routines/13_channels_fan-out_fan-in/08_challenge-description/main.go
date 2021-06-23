package main

import (
	"fmt"
	"sync"
)

var wordId int
var publishId int

func main() {
	in := gen()
	x := facOut(in, 10)
	//fmt.Printf("%T \n", x)
	//fmt.Println("*******************", len(x))
	//for _, v := range x {
	//	fmt.Println("********", <-v)
	//}
	//

	for v := range merge(x...) {
		fmt.Println(v)
	}

}

func gen() <-chan int {
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

//Note that a chan is returned directly here
func facOut(in <-chan int, n int) []<-chan int {
	var out []<-chan int
	for i := 1; i <= n; i++ {
		//Note that a chan is returned directly here, No need for goroutine
		out = append(out, factorial(in))
	}
	return out
}

func factorial(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range c {
			out <- fact(v)
		}
		close(out)
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

func merge(c ...<-chan int) <-chan int {
	fmt.Println("merge %T", c)
	var wg sync.WaitGroup
	out := make(chan int)
	outFunc := func(c <-chan int) {
		for v := range c {
			out <- v
		}
		wg.Done()
	}
	//The add here must be placed outside
	wg.Add(len(c))

	for _, v := range c {
		go outFunc(v)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
