package main

import "fmt"

var wordId int
var publishId int

func main() {
	in := gen()
	for v := range fact(in) {
		fmt.Println(v)
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

// 多个函数从同一通道读取数据，直到该通道关闭。
// 将工作分布在多个函数（10个goroutines）中，这些函数都从in读取。
