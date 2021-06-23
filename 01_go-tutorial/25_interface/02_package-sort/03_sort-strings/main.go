package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"alice", "bob", "Ai", "Jenny"}
	sort.Strings(s)
	fmt.Println(s)
}
