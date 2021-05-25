package main

import "fmt"

func main() {
	var x [256]string
	fmt.Println("x is ", x)
	fmt.Println("len(x) is ", len(x))
	fmt.Println("x[42] is ", x[42])

	for i := 0; i < 256; i++ {
		x[i] = string(i)
	}
	fmt.Println("for loop x[42] is ", x[42])

	for i, v := range x {
		fmt.Printf("%v -- %T -- %b \n", v, v, []byte(v))
		if i > 50 {
			break
		}
	}

}
