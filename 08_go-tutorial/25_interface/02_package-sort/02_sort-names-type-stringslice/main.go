package main

import (
	"fmt"
	"sort"
)

func main() {
	studyGroup := []string{"alice", "bob", "Ai", "Jenny"}
	fmt.Println(studyGroup)
	//type conversation
	sort.Sort(sort.StringSlice(studyGroup))
	fmt.Println(studyGroup)

}
