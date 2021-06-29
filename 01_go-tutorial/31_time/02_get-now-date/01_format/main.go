package main

import (
	"fmt"
	"time"
)

var (
	TimeFormat   = "20060102"
	TimeFormatV2 = "2006-01-02"
)

func main() {
	fmt.Println(time.Now().Format(TimeFormat))
	fmt.Println(time.Now().Format(TimeFormatV2))
}
