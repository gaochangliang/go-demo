package main

import "fmt"

func main() {
	switch "kobe" {
	case "kobe", "Jordan":
		fmt.Println("basketball")
	case "CR", "MX":
		fmt.Println("football")
	}
}

/*
  no default fallthrough
  fallthrough is optional
  -- you can specify fallthrough by explicitly stating it
  -- break isn't needed like in other languages
*/
