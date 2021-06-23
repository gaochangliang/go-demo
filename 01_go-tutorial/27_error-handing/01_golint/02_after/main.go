package main

import "fmt"

func main() {
	x := 1
	str := evalInt(x)
	fmt.Println(str)
}

func evalInt(x int) string {
	if x > 10 {
		return fmt.Sprintf("x is grater than 10")
	}
	return fmt.Sprintf("x is less than 10")

}
