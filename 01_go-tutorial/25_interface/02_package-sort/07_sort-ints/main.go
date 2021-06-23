package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{10, 2, 4, 6, 8, 9, 3, 33}
	sort.Ints(n)
	fmt.Println("sort", n)
}
