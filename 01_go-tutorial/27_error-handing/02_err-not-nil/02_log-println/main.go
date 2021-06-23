package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("a.txt")
	if err != nil {
		log.Println("err happened", err)
	}

}

// Println formats using the default formats for its operands and writes to standard output.

// Println使用其操作数的默认格式进行格式化，并写到标准输出。

/*

包log实现了一个简单的日志包......写入标准错误并打印每条日志信息的日期和时间......

Fatal函数在写入日志信息后调用os.Exit(1)......
Panic函数在写入日志信息后调用panic。
Println 调用 Output 来打印到标准的记录器。参数以 fmt.Println 的方式处理

*/
