package gee

import (
	"net/http"
)

type Engine struct {
	router *router
}

type handlerFunc func(*Context)

func New() *Engine {
	return &Engine{router: newRouter()}
}

//注册路由
func (engine *Engine) GET(pattern string, handler handlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) addRoute(method, pattern string, handler handlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//context接受了请求参数，那么必然处理的逻辑肯定要写在context里面
	c := newContext(w, req)
	engine.router.handle(c)
}

//注意这里的第二个参数是传的调用者本身
func (engine *Engine) Run(pattern string) error {
	return http.ListenAndServe(pattern, engine)
}
