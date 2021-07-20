package main

import (
	"flag"
	"fmt"
)

var (
	name string
	age  int
	fish bool //Whether you can fish
)

func main() {

	//此处的第二个参数定义了命令行需要指定的参数
	//注意这里是指针绑定
	flag.StringVar(&name, "name", "bar", "input the name")
	flag.IntVar(&age, "age", 28, "input the age")

	flag.Parse()

	fmt.Println("name", name)
	fmt.Println("age", age)

}

/*

Go provides a flag package supporting basic command-line flag parsing.


exec:
go build
main --name="kobe" --age=41


*/

/*

you provide a flag that wasn’t specified to the flag package,
the program will print an error message and show the help text again.

main --name1="kobe"
flag provided but not defined: -name
Usage of F:\CodeDemo\go-demo\01_go-tutorial\42_pkg-self\flag\main.exe:
-a int
input the age (default 28)
-n string
input the name (default "bar")

*/
