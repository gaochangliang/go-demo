package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int

func main() {
	wg.Add(2)
	go incrementer("foo")
	go incrementer("bar")
	wg.Wait()
}

func incrementer(s string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= 20; i++ {
		x := counter
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		x++
		counter = x
		fmt.Println(" x", x, " s ", s, "counter", counter)
	}
	wg.Done()
}

// go run -race main.go
// vs
// go run main.go
