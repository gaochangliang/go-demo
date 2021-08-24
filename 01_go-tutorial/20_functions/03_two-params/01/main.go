package main

import "fmt"

func main() {
	greet("kobe", "kobe Bryant")
}

//If it is not a separate print, but designed to some calculations and other operations,
//need to consider the parameter is empty will report an error case
func greet(fname string, lname string) {
	fmt.Printf("fname is %s lname is %s\n", fname, lname)
}
