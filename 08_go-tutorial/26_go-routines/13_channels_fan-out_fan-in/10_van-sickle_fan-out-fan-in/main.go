package main

import "fmt"

func main() {
	in := gen()
	out := make(chan int)
	factOut(in, 10, out)

	go func() {
		for v := range out {
			fmt.Println(v)
		}
	}()

	var a string
	fmt.Scanln(&a)

}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
	}()
	return out
}

//这里n的作用只是协程的数量
//相当于有n个协程在把数据从in读到n中
func factOut(in <-chan int, n int, out chan int) {
	for i := 0; i < n; i++ {
		factorial(in, out)
	}
}

func factorial(in <-chan int, out chan<- int) {
	go func() {
		for n := range in {
			out <- fact(n)
		}
	}()
}

func fact(n int) int {
	total := 1
	for i := 2; i <= n; i++ {
		total *= i
	}
	return total
}
