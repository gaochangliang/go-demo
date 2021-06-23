package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	xs := visit(nums, printNumber)
	fmt.Println(xs)
}

func visit(numbers []int, callback func(int) bool) []int {
	var xs []int
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	return xs
}

func printNumber(number int) bool {
	return number > 2
}
