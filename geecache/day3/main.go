package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	var str = `mH9qbu1tP3uG1Uvn66qL8Ad\/S7uxd6Ap6FQ4SWJ4t\/8=`
	buffer,_ := base64.StdEncoding.DecodeString(str)
	fmt.Println(hex.EncodeToString(buffer))
}
