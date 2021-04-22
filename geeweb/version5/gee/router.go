package gee

import "strings"

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

type HandlerFunc func(*Context)

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

}

/*
/admin/get return ["admin","get"]
/admin  return ["admin"]
*   return ["*"]
*/

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item == "*" { //support only *
				break
			}
		}
	}
	return parts
}
