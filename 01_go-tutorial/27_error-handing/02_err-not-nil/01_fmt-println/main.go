package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("a.txt")
	if err != nil {
		fmt.Println("err happen", err)
	}

}

// Println formats using the default formats for its operands and writes to standard output.

// Println使用其操作数的默认格式进行格式化，并写到标准输出。
