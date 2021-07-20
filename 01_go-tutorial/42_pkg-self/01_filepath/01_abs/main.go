package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Abs("."))
	fmt.Println(filepath.Abs("/abc"))
	fmt.Println(filepath.Abs("abc"))
	fmt.Println(filepath.Abs(".."))
}

//abs监测传入的当前路径，如果是决定路径直接返回,
//不是绝对路径直接返回当前路径 + 传入路径
