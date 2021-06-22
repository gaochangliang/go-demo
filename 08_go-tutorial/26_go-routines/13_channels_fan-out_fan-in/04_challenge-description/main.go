package main

import (
	"fmt"
	"time"
)

var wordId int
var publishId int

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
	publishId++
	thisId := wordId
	dataId := 0
	for {
		dataId++
		fmt.Printf("publish id %d is publishing data", thisId)
		data := fmt.Sprintf("data from publisher wordid- %d data- %d ", wordId, dataId)
		out <- data
	}
}

func workerProcess(in <-chan string) {
	wordId++
	thisId := wordId
	for {
		fmt.Printf("%d: waiting for input...\n", thisId)
		input := <-in
		fmt.Printf("%d: input is: %s\n", thisId, input)
	}

}
