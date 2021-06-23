package main

import "fmt"

func main() {
	input := map[int]int{}
	input[1] = 123
	input[3] = 12345
	input[5] = 1234567

	c := make(chan int)
	go func() {
		for _, v := range input {
			c <- v
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}

//有一个map，需要异步读取map的数据到协程里面，自己遍历完成后关闭通道
//然后遍历完成,
