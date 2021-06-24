package main

import "fmt"

type mySentence string

func main() {
	var myS mySentence = "hello world"
	fmt.Println("string", myS)
	fmt.Printf("%T\n", myS)
	fmt.Printf("%T\n", myS.(string))
}

//不是接口类型类型断言失败
//it is not interface ,assert failed
