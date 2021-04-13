package main

import (
	"fmt"
	"path"
)

func main() {
	filePath := "/"
	relPath := "/"
	finalPath := path.Join(filePath,relPath)
	fmt.Println(finalPath)
}
