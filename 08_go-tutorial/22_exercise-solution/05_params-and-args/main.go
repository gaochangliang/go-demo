package main

import "fmt"

func main() {
	foo(1, 2)
	foo(1, 2, 3)

	var aSlice = []int{100, 200, 300}
	foo(aSlice...)

	foo()

}

func foo(numbers ...int) {
	fmt.Println(numbers)
}
