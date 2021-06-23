package main

import "fmt"

func main() {
	name := "kobe"
	fmt.Println("name length is", len(name))
	switch {
	case name == "kobe":
		fmt.Println("basketball")
	case len(name) == 4:
		fmt.Println("basketball")
	case name == "CR":
		fmt.Println("football")
	}
}

/*


/*
  expression not needed
  -- if no expression provided, go checks for the first case that evals to true
  -- makes the switch operate like if/if else/else
  cases can be expressions

 we can senn There are two matches, but they end as soon as they are matched
*/
