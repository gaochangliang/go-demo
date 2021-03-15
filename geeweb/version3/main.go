package main

import (
	"fmt"
	"github/gee/version3/gee"
	"net/http"
)

func main() {
	r := gee.New()

	//r.Get只是我们把对应的函数传给了处理的r保存起来,真正在处理的地方还是Run函数
	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path=%q", req.URL.Path)
	})

	r.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	r.Run(":8080")
}
