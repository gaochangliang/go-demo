package router

import (
	v1 "gin/05/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup) {
	initRouter := r.Group("init")
	{
		initRouter.POST("initdb", v1.InitDB)
	}
}
