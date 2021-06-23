package main

import "fmt"

func main() {
	greet("kobe", "kobe Bryant")
}

func greet(fname string, lname string) {
	fmt.Printf("fname is %s lname is %s\n", fname, lname)
}
