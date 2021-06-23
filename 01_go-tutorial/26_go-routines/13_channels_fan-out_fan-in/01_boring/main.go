package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("alice"), boring("bob"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func boring(word string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %v", word, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(chan1, chan2 <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for {
			out <- <-chan1
		}

	}()

	go func() {
		//notes the loop here
		for {
			out <- <-chan2
		}

	}()
	return out
}
