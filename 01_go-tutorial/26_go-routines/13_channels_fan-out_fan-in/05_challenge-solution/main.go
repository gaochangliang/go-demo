package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var wordId int64
var publishId int64

func main() {
	input := make(chan string)
	go workerProcess(input)
	go workerProcess(input)
	go workerProcess(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
	time.Sleep(1 * time.Second)
}

func publisher(out chan string) {
	atomic.AddInt64(&publishId, 1)
	thisId := atomic.LoadInt64(&publishId)
	dataId := 0
	for {
		dataId++
		fmt.Printf("publish id %d is publishing data", thisId)
		data := fmt.Sprintf("data from publisher wordid- %d data- %d ", wordId, dataId)
		out <- data
	}
}

func workerProcess(in <-chan string) {
	atomic.AddInt64(&wordId, 1)
	thisId := atomic.LoadInt64(&wordId)
	for {
		fmt.Printf("%d: waiting for input...\n", thisId)
		input := <-in
		fmt.Printf("%d: input is: %s\n", thisId, input)
	}

}
