## 1 day1-http-base

介绍简单的原生的 `http` 服务的搭建

**定义路由**

```go
http.HandleFunc("/", indexHandler)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}	
```

**启动服务**

`nil` 默认采用默认标准库中的实例处理方式，`nil` 这个参数也是 `web` 框架的入口

```go
log.Fatal(http.ListenAndServe(":8080", nil))
```

## 2 初次接触geeinit

默认的处理器实现了如下接口

```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

因此想要替代原生的处理器方式必须实现这个接口，这样第二个参数传入我们自己的处理器

```go
type Engine struct {

}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	case "/hello":
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}
```

**启动服务**

```go
engine := new(Engine)
log.Fatal(http.ListenAndServe(":9999", engine))
```



## 3  优化路由

观察上面的 `ServeHTTP` 所有的处理器方式都在这一个函数中，不易扩展，我们可以将路由作为 `engine` 的一个属性，是一个 `map`结构，`key` 对应路由，`value` 对应处理函数，同时提供增加路由的方式  `GET` 和 `POST`，

```go
type handlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]handlerFunc    //key 代表注册路由 value代表对应路由处理函数
}
```

**注册路由到`engine`中**

```go
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
```

**路由注册后这样就有新的处理方式`ServeHTTP`了**

```go
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}
```

**启动服务**

```go
g := gee.New()
g.Get("/", func(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
})
g.Run(":8080")
```

