package main

import (
	"fmt"
	"path"
	"strings"
)

func main() {
	var imgName = "zhaowei.jpg"
	GetSuffixExt(imgName)
}

func GetSuffixExt(name string) {
	ext := path.Ext(name)
	//name去掉后缀ext,没有包含后缀ext直接返回原值
	suffix := strings.TrimSuffix(name, ext)

	fmt.Println("suffix", suffix)
	fmt.Println("ext", ext)
}
