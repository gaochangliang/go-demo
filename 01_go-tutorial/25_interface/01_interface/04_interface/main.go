package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (z square) area() float64 {
	return z.side * z.side
}

type shape interface {
	area() float64
}

func totalArea(shapes ...shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func info(z shape) {
	fmt.Println("info", z)
	fmt.Println("info", z.area())
}

func main() {
	s := square{10}
	c := circle{100}
	info(s)
	info(c)
	fmt.Println("Total area: ", totalArea(s, c))
}
