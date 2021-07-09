package routers

import (
	v1 "example/01_qrcode/pkg/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	}

	return r
}
