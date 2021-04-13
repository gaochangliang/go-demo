//Get the current version of go
package main

import (
	"runtime"
	"strconv"
	"strings"
)

func main() {
	getMinVer(runtime.Version())
}

func getMinVer(v string) (uint64, error) {
	firstIndex := strings.IndexByte(v, '.')
	lastIndex := strings.LastIndexByte(v, '.')
	if firstIndex == lastIndex {
		return strconv.ParseUint(v[firstIndex+1:], 10, 64)
	}
	return strconv.ParseUint(v[firstIndex+1:lastIndex], 10, 64)
}
