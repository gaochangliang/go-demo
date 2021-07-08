package main

import "example/01_create_dir/pkg/file"

const (
	imagePath = "image/"
)

func main() {
	file.MustOpen(imagePath)
}
