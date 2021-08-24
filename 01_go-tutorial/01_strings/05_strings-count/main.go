package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "abcda"
	fmt.Println(strings.Count(str,""))	//判断子串在源串中的数量，如果子串为空，则长度为源串的长度+1。
	fmt.Println(strings.Count(str,"a"))	//
}
