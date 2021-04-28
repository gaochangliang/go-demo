package main

import "fmt"

const suffix = "\n"

func main() {
	var s = "\na1cb\n"
	fmt.Println(HasSuffix(s, suffix))
}

// HasSuffix tests whether the string s ends with suffix.
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}
