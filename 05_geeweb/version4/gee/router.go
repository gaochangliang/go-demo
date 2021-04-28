package gee

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]handlerFunc
}

func newRouter() *router{
	return &router{handlers: map[string]handlerFunc{}}
}

func (router *router) addRoute(method, pattern string, handler handlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	router.handlers[key] = handler
}

func (router *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	log.Println(key)
	if handler, ok := router.handlers[key];ok {
		handler(c)
	}else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
