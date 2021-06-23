package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64

func main() {
	wg.Add(2)
	go incrementer("foo")
	go incrementer("bar")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementer(s string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		atomic.AddInt64(&counter, 1)
		//Because this line is not locked, there will be competition
		//atoom
		fmt.Println(" s ", s, "counter", counter)

	}
	wg.Done()
}

// go run -race main.go
// vs
// go run main.go
