package main

import "fmt"

func getSum() func() int {
	sum := 1
	return func() int {
		sum++
		return sum
	}
}

func main() {
	//这里每个函数都是独立的不共享的，但是这是一个闭包，sum和返回的函数形成闭包，多次调用
	//同一个函数sum共享
	s1 := getSum()
	s2 := getSum()
	fmt.Println("s1", s1())
	fmt.Println("s2", s2())

	fmt.Println("s1s1", s1())
	fmt.Println("s2s2", s2())
}
