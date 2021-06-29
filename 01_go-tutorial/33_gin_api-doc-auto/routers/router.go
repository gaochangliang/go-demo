package routers

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery()) //panic默认造成崩溃退出，recovery处理这个异常然后返回http code 500。
	return r
}
