package initalize

import (
	"gin/05/middleware"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var r = gin.Default()
	r.Use(middleware.Cors())

	publicGroup := r.Group("")
	{
		r.InitInitRouter(publicGroup) // 自动初始化相关
	}
}
