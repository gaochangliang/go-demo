package main

import "fmt"

const (
	testStr  = "kobe"
	testStr1 = "abcde234241"
)

func main() {
	fmt.Println("testStr", reverseTwo(testStr))
	fmt.Println("testStr2", reverseTwo(testStr1))
}

func reverseTwo(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
