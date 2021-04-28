package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "/"
	vs := strings.Split(str, "/")
	fmt.Printf("str split %v 	length %d\n",vs,len(vs))


	str1 := "/admin"
	vs1 := strings.Split(str1, "/")
	fmt.Printf("str1 split %v length %d\n",vs1,len(vs1))


	str2 := "/admin/login"
	vs2 := strings.Split(str2, "/")
	fmt.Printf("str2 split %v length %d\n",vs2,len(vs2))

	str3 := "*"
	vs3 := strings.Split(str3, "/")
	fmt.Printf("str3 split %v length %d\n",vs3,len(vs3))
}
