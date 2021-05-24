package main

import "fmt"

func main() {
	fmt.Println(42)               //decimal
	fmt.Printf("%d-%b\n", 42, 42) //binary

	/*
		utf-8
		%d 	decimal
		%b 	binary
		%x 	hexadecimal Lowercase
		%q
	*/

	/*
		Escape characters
		\n  Line feed
		\r  Current position moved to the beginning of this line
		\t  Horizontal tab (HT) 水平制表符
	*/

	for i := 60; i < 122; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
