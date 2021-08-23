package main

import (
	"fmt"
)

//函数的外面不能使用 := 定义变量
//这种变量称为包内部变量
const N = "kobe"

func main() {
	const a = 41
	fmt.Println("N -", N)
}
