package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	//notes Note that this is concurrent execution
	//The order of the results is variable
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i <= 45; i++ {
		fmt.Println("foo i", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i <= 45; i++ {
		fmt.Println("bar i", i)
	}
	wg.Done()
}
