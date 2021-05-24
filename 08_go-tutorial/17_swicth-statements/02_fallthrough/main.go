package main

import "fmt"

func main() {
	switch "kobe" {
	case "kobe":
		fmt.Println("good man kobe")
		fallthrough
	case "baoluo":
		fmt.Println("good man baoluo")
		fallthrough
	case "wade":
		fmt.Println("good man wade")
	case "jordan":
		fmt.Println("good man jordan")
	}
}

/*
  no default fallthrough
  fallthrough is optional
  -- you can specify fallthrough by explicitly stating it
  -- break isn't needed like in other languages
*/
