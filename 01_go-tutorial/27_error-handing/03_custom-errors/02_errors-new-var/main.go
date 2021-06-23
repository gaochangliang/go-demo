package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrNorteMathErr = errors.New("norgate math: square root of negative number")

func main() {
	fmt.Printf("ErrNorteMathErr type %T\n", ErrNorteMathErr)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNorteMathErr
	}
	return 42, nil
}
