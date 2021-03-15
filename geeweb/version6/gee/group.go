package gee

import "log"

type RouterGroup struct {
	prefix      string
	middleWares []handlerFunc //支持的中间件
	engine      *Engine       //
}

func (group *RouterGroup) addRoute(method, pattern string, handler handlerFunc) {
	pattern = group.prefix + pattern
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler handlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler handlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}
