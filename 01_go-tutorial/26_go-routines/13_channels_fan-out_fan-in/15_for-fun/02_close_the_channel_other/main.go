package main

import "fmt"

func main() {
	input := map[int]int{}
	input[1] = 123
	input[3] = 12345
	input[5] = 1234567

	c := make(chan int)
	ch2 := make(chan int)

	go func() {
		var i int
		for n := range ch2 {
			i += n
			if i == len(input) {
				close(c)
			}
		}
	}()

	go func() {
		for _, v := range input {
			c <- v
			ch2 <- 1
		}
	}()

	for v := range c {
		fmt.Println(v)
	}

	close(ch2)
}

//通过一个goroutine控制另一个协程数据的读取和终止

//有一个map，需要异步读取map的数据到协程里面，并不是自己关闭通道，而是自己每读取一个数据后，给另一个通道2写数据
//通道2有终止通道1读取数据的条件，i == len(input),最后关闭通道2
//通道2完全就是为通道1服务的
