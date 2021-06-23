package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	fmt.Println("cpu number", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i <= 45; i++ {
		fmt.Println("foo i", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}

func bar() {
	for i := 0; i <= 45; i++ {
		fmt.Println("bar i", i)
		time.Sleep(20 * time.Millisecond)
	}
	wg.Done()
}
