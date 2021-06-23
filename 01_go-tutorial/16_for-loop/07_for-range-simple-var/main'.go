package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3, 4, 5}
	for k, v := range nums {
		fmt.Println("kv", k, "-", v)
	}

	//Single variable loops is index
	for k := range nums {
		fmt.Println("k", k)
	}
}
