package gee

import (
	"fmt"
	"net/http"
	"strings"
)

type router struct {
	handlers map[string]handlerFunc //key请求方式+路由
	roots    map[string]*node       //key是请求方式 e.g. get、post
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]handlerFunc),
		roots:    make(map[string]*node),
	}
}

func (r *router) addRoute(method string, pattern string, handler handlerFunc) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

func (router *router) handle(c *Context) {
	fmt.Printf("c.Method[%v] c.path[%v]\n", c.Method, c.Path)
	n, params := router.getRoute(c.Method, c.Path)

	//return
	if n != nil {
		c.params = params
		key := c.Method + "-" + n.pattern
		c.handlers = append(c.handlers, router.handlers[key])
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
	c.Next()
}

//去掉路由的/,返回一个切片字符串
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/") //返回一个切片
	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	//
	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
