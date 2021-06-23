package main

import (
	"fmt"
	"log"
)

type NorgateMathError struct {
	lat, long string
	err       error
}

func (n *NorgateMathError) Error() string {
	return fmt.Sprintf("norgate math: square root of negative number:%v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		norgateErr := fmt.Errorf("norgate math: square root of negative number:%v", f)
		return 0, &NorgateMathError{
			"-51", "-52", norgateErr,
		}
	}
	return 42, nil
}
