package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

type byAge []person

func (b byAge) Less(i, j int) bool { return b[i].Age < b[j].Age }
func (b byAge) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byAge) Len() int           { return len(b) }

func main() {
	people := []person{
		{Name: "kobe", Age: 41},
		{Name: "curry", Age: 16},
		{Name: "jordan", Age: 100},
	}
	fmt.Println(people[0])
	fmt.Println(people)
	sort.Sort(byAge(people))
	fmt.Println(people)
}
