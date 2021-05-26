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

## 4 context

对Web服务来说，无非是根据请求`*http.Request`，构造响应`http.ResponseWriter`，之前的所有的处理函数都是这个形式` func(w http.ResponseWriter, r *http.Request)` ,这种方式设置响应头的话，所有的函数都要一一设置，需要进行封装,返回的信息通常有如下几种，`JSON,string`等等

**设置一个新的结构保存请求响应信息 `w http.ResponseWriter, r *http.Request`**

```
type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request
	
	Path   string
	Method string
	StatusCode int
}
```

## 修改路由处理函数

之前的路由处理函数

```go
type handlerFunc func(w http.ResponseWriter, r *http.Request)
```

修改成如下

```go
type HandlerFunc func(*Context)
```

## context 获取请求参数函数

别忘了context包含了请求信心以及处理信息，对于请求信心来说，通常需要获取用户传过来的参数

因此此处封装两个回去用户传参的函数

`GET请求`

```go
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}
```

`POST请求`

```go
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}
```

## context处理响应函数

context根据之前的描述，需要封装像一个，比如状态码，响应头等等

```go
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}
```

设置通用的处理函数，返回字符串，`JSON`等函数

```go
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
```

路由基本保持不变

## 启动服务

```go
r := gee.New()
r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
})

r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
})

r.Run(":9999")
```

