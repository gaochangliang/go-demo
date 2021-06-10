package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) float64 {
	return s.area()
}

func main() {
	c := circle{10}
	//error
	cArea := info(c)
	fmt.Println("area", cArea)
}
