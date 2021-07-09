package main

import "example/02_create_dir/pkg/file"

const (
	imagePath = "image/"
	fileName  = "a.jpg"
)

func main() {
	//创建文件
	file.MustOpen(fileName, imagePath)
}
