package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("a.txt")
	if err != nil {
		log.Fatalln("err", err)
	}
}

//包log实现了一个简单的日志包......写入标准错误并打印每条日志信息的日期和时间......
//Fatal函数在写入日志信息后调用os.Exit(1)......Panic函数在写入日志信息后调用panic。
// Fatalln相当于Println()后调用os.Exit(1)。
