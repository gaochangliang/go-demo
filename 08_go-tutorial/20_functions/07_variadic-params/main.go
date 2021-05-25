package main

import "fmt"

func main() {
	n := average(100, 200, 300, 300, 400)
	fmt.Println("n", n)
}

func average(sf ...float64) float64 {
	fmt.Println(" sf ", sf)
	fmt.Printf(" %T \n", sf)
	var total float64
	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))
}

//variadic 可变参数
