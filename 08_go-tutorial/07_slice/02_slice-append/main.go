package main

import "math"

const abortIndex int8 = math.MaxInt8 / 2

type Context struct {
}

type RouterGroup struct {
	Handlers HandlersChain
	basePath string
}

type HandlerFunc func(*Context)

type HandlersChain []HandlerFunc

type Engine struct {
	routeGroup RouterGroup
	allNoRoute HandlersChain
}

/*
将routeGroup和接收的HandlersChain赋值给allNoRoute

注意创建新的路由的写法
make写法
copy用法
*/

func (group *RouterGroup) combineHandlers(handlers HandlersChain) HandlersChain {
	finalSize := len(group.Handlers) + len(handlers)
	if finalSize >= int(abortIndex) {
		panic("too many handlers")
	}
	mergedHandlers := make(HandlersChain, finalSize)
	copy(mergedHandlers, group.Handlers)
	copy(mergedHandlers[len(group.Handlers):], handlers)
	return mergedHandlers
}
