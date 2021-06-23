package main

import "fmt"

func main() {
	greet("kobe", "kobe Bryant")
}

func greet(fname, lname string) {
	fmt.Printf("name is %s lname is %s\n", fname, lname)
}
