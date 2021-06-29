package main

import (
	"fmt"
	"os"
)

const logName = "logs/"

func main() {
	mkdir()
}

func mkdir() {
	//返回与当前目录对应的根路径名
	dir, _ := os.Getwd()
	fmt.Println("dir", dir)
	err := os.MkdirAll(dir+"/"+logName, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
