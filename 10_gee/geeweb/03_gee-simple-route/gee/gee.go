package gee

import (
	"fmt"
	"net/http"
)

//如何定义函数类型
type handlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]handlerFunc
}

func New() *Engine {
	return &Engine{
		router: make(map[string]handlerFunc),
	}
}

func (e *Engine) Get(router string, handlerF handlerFunc) {
	e.addRoute("GET", router, handlerF)
}

func (e *Engine) POST(router string, handlerF handlerFunc) {
	e.addRoute("POST", router, handlerF)
}

func (e *Engine) addRoute(method, pattern string, handlerF handlerFunc) {
	key := method + "-" + pattern
	e.router[key] = handlerF
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}

// Run defines the method to start a http server
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}
