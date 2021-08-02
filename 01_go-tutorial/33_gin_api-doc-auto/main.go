package main

import (
	"fmt"
	_ "gin-doc/docs"
	"gin-doc/routers"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"time"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample api celler api.
// @termsOfService http://swagger.io/terms/
func main() {
	router := routers.InitRouter()
	//注意这里的类型一定要是int,字符串导致服务不能启动
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	var httpPort = 81
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", httpPort),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		/*
			3 << 2，则是将数字3左移2位
			计算过程：
			3 << 2 首先把3转换为二进制数字0000 0011，然后把该数字高位(左侧)的两个零移出，其他的数字都朝左平移2位，最后在低位(右侧)的两个空位补零。
			则得到的最终结果是0000 1100，则转换为十进制是12。
		*/
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("[info] start http api listening %d", httpPort)

	_ = server.ListenAndServe()

}
