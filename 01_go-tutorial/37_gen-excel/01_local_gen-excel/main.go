package main

import (
	v1 "example/gen-excel/pkg/api/v1"
	"example/gen-excel/pkg/setting"
	"fmt"
)

func init() {
	setting.Setup()
}

func main() {
	var tag = &v1.Tag{}
	s, err := tag.Export()
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("s", s)
}
