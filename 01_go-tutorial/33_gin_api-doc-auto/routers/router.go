package routers

import (
	"gin-doc/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())   //使用gin的日志
	r.Use(gin.Recovery()) //panic默认造成崩溃退出，recovery处理这个异常然后返回http code 500
	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
	}
	return r

}
