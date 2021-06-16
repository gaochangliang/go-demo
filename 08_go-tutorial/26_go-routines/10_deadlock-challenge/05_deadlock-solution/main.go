package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	//A closed chan does not cause a deadlock, but it is an ongoing loop because, if a closed chan takes a value, it can take a zero value and will not end
	//关闭的chan虽然不会造成deadlock，但是这是一个一直循环，因为，关闭的chan取值的话可以取到零值，不会结束
	for v := range c {
		fmt.Println(v)
	}

}

// This prints the number, but we have received this error:
// fatal error: all goroutines are asleep - deadlock!
// Where is the deadlock?
// Why are all goroutines asleep?
// How can we fix this?
