package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%41 == 0 {
			fmt.Println(i, " -- kobe ")
		} else if i%23 == 0 {
			fmt.Println(i, " -- Jordan ")
		} else if i%30 == 0 {
			fmt.Println(i, " -- thirty")
		} else {
			fmt.Println(i)
		}
	}
}

//use numOne not num1
