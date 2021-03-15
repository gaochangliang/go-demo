package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

//Context是保存用户传过来的参数内容
type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	StatusCode int
	Method     string
	Path       string
	params     map[string]string
	searchCode int
}

//前面说过Context是保存用户传过来的参数内容，那么一定会有这两个参数
func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	fmt.Println("name", c.Req.URL.Query().Get(key))
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHandler("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHandler("Content-Type", "text/plain")
	c.Status(code)
	//NewEncoder创建一个将数据写入w的*Encode
	encoder := json.NewEncoder(c.Writer)
	//方法Encode将obj的json编码写入输出流，并会写入一个换行符
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHandler("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

/* 设置响应消息的头部
Header().Set 然后WriteHeader() 最后是Write()
在 WriteHeader() 后调用 Header().Set 是不会生效的
*/

func (c *Context) SetHandler(key, val string) {
	c.Writer.Header().Set(key, val)
}
