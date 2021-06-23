package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{10, 2, 4, 6, 8, 9, 3, 33}
	fmt.Println("start", n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println("sort", n)
}
