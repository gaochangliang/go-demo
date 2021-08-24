package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}

	sliOne := slice1[0:5]	//slice[start:end] end最多能取长度5
	fmt.Println(" 0 - 5 ", sliOne)

	sliTwo := slice1[0:6]
	fmt.Println(sliTwo)
}

/*

because it is left-closed right-open [a:b] a and b represent the index position, the actual cut index is taken as a - b-1

The actual index of slice1 can be in the range [0:4].

slice1[0:5] gets 0-4

*/
