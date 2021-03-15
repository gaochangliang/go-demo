package main

import (
	"fmt"
	"log"
	"net/http"
)

//type Handler interface {
//	ServeHTTP(ResponseWriter, *Request)
//}

//只要传入任何实现了 ServerHTTP 接口的实例，所有的HTTP请求，就都交给了该实例处理
type Engine struct{}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "url.path = %v", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND:%s\n", req.URL.Path)
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":80", engine))
}
