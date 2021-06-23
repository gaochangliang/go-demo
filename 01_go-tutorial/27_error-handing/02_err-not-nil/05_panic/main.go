package main

import "os"

func main() {
	_, err := os.Open("a.txt")
	if err != nil {
		panic(err)
	}
}
