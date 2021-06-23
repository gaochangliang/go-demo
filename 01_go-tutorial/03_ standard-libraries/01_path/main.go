package main

import "path"

func main() {
	basePath := "/"
	relativePath := "/"
	finalPath := path.Join(basePath, relativePath)
	appendSlash := lastChar(relativePath) == '/' && lastChar(finalPath) != '/'
	if appendSlash {
		finalPath = finalPath + "/"
	}
}

func lastChar(str string) uint8 {
	if str == "" {
		panic("The length of the string can't be 0")
	}
	return str[len(str)-1]
}
