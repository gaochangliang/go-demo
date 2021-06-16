package main

import "fmt"

func main() {
	n := factorial(4)
	fmt.Println("total", n)
}

func factorial(n int) int {
	total := 1
	for i := 1; i <= n; i++ {
		total *= i
	}
	return total
}
