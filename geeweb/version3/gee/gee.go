package gee

import (
	"fmt"
	"net/http"
)

type handleFunc func(w http.ResponseWriter, req *http.Request)

type Engine struct {
	router map[string]handleFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]handleFunc)}
}


func (engine *Engine) addRoute(method string, pattern string, handler handleFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

//
func (engine *Engine) Get(pattern string, handler handleFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler handleFunc) {
	engine.addRoute("POST", pattern, handler)
}
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND %s\n", req.URL)
	}
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

