package main

import "fmt"

func main() {
	var aName, bName string = "alice", "bob"
	i, j, k := 1, 2, 3
	item, price := "iphone", 6000
	fmt.Println(aName + bName)
	fmt.Println(i + j + k)
	fmt.Println(item, "-", price)
}
