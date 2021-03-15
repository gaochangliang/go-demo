package main

import (
	"fmt"
	"github/gee/version7/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.Use(gee.Logger()) //全局中间件
	r.GET("/", func(c *gee.Context) {
		fmt.Println("1111")
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	v2 := r.Group("/v2")
	v2.Use(gee.OnlyForV2())
	{
		v2.GET("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v2.GET("/hello", func(c *gee.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	r.Run(":8080")
}
