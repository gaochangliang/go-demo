package main

import "fmt"

type person struct {
	names []string
}

func main() {
	var p person
	p.addName("sunwukong", "sunxingzhe")
	fmt.Println("person names", p.names)

	nameStorage := make([]string, 2)
	nameStorage[0] = "tangsanzang"
	nameStorage[1] = "shawujing"
	p.addName(nameStorage...)
	fmt.Println("person names", p.names)
}

func (p *person) addName(names ...string) {
	for _, v := range names {
		p.names = append(p.names, v)
	}
}
