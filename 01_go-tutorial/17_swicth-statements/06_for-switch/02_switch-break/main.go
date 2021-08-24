package main

import "fmt"

func main() {
	j := 1
	for i := 0; i < 4; i++ {
		switch j {
		case 1:
			fmt.Println("1")
			j++
			break			//其实每个case后面默认会有break，这个break的范围只在switch中
		case 2:
			//结束本次循环执行下一次循环,下面的2不会打印
			fmt.Println("2")
			continue
		}
	}
}
