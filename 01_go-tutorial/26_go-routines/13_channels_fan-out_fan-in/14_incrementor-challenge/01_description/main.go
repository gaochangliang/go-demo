package main

import "fmt"

func main() {
	c := incrementor(2)
	for v := range c {
		fmt.Println(v)
	}
}

func incrementor(n int) chan string {
	c := make(chan string)
	out := make(chan bool)
	for i := 0; i <= n; i++ {
		go func(number int) {
			c <- fmt.Sprintf("incrementor %v", number)
			out <- true
		}(i)
	}

	//主要通过数字number控制所有的协程结束，然后close掉
	go func() {
		for i := 1; i <= n; i++ {
			<-out
		}
		close(c)
	}()
	return c
}

//The number number is mainly used to control close() after the completion of the concurrent process.
