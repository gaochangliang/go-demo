package main

import "fmt"

func main() {
	fmt.Println(greet("kobe", "bryant"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}

/*
IMPORTANT
Avoid using named returns.


Occasionally named returns are useful. Read this article for more information:
https://www.goinggo.net/2013/10/functions-and-naked-returns-in-go.html
*/
