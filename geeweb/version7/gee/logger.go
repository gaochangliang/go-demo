package gee

import (
	"log"
	"time"
)

func Logger() handlerFunc {
	return func(context *Context) {
		t := time.Now()
		//这里使用next表示先调用其他的中间件,最后再输出下面这条日志
		context.Next()
		log.Printf("[%d] %s in %v", context.StatusCode, context.Req.RequestURI, time.Since(t))
	}
}

func OnlyForV2() handlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		//c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
