package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	visit(nums, printNumber)
}

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func printNumber(number int) {
	fmt.Println(number)
}

// callback: passing a func as an argument
