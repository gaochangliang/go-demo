package main

import "fmt"

func main() {
	data := []float64{100, 200, 300, 300, 400}
	n := average(data...)
	fmt.Println("n", n)
}

func average(sf ...float64) float64 {
	var total float64
	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))
}
