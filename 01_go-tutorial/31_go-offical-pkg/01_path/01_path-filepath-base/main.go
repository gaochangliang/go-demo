package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	var file1 = "logs/2021.log"
	var file2 = "logs/2021/"
	var file3 = "logs/2021"
	var file4 = ""

	fmt.Println("file1", filepath.Base(file1))
	fmt.Println("file2", filepath.Base(file2))
	fmt.Println("file3", filepath.Base(file3))
	fmt.Println("file4", filepath.Base(file4))
}
