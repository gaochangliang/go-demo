/*
base64字符串转hex
hex字符串转base64
*/
package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func Base64ToHex(str string) string {
	buffer, _ := base64.StdEncoding.DecodeString(str)
	return hex.EncodeToString(buffer)
}

func HexToBase64(str string) string {
	buffer, _ := hex.DecodeString(str)
	return base64.StdEncoding.EncodeToString(buffer)
}

func main() {
	fmt.Println(1)
	fmt.Println(Base64ToHex(`+BW1qcip9420DrZHF6Cas88xGvkSSrDPazS/ObYqXMI=`))
}