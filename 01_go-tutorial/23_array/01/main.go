package main

import "fmt"

func main() {
	var x [58]int
	fmt.Println("x is ", x)
	fmt.Println("len(x) is ", len(x))

	fmt.Println("x[42] is ", x[42])
	x[42] = 4200
	fmt.Println("update x[]42 is ", x[42])

}
