package main

import "fmt"

func main() {
	j := 1
	for i := 0; i < 4; i++ {
		switch j {
		case 1:
			fmt.Println("1")
			j++
		case 2:
			//结束本次循环执行下一次循环,下面的2不会打印
			continue
			fmt.Println("2")
		}
	}
}
