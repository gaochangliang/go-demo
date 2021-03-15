package main

import (
	"github/gee/version4/gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.GET("/",func(c *gee.Context){
		c.JSON(http.StatusOK, gee.H{
			"username": c.Query("username"),
			"password": c.Query("password"),
		})
	})
	//重点是http.ListenAndServe(addr, engine)中的engine，engine的ServeHTTP处理函数才是处理请求的逻辑
	r.Run(":9999")
}

