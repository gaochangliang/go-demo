package routers

import (
	_ "02_go-gin-example/gin-blog/docs"
	"02_go-gin-example/gin-blog/middleware/jwt"
	"02_go-gin-example/gin-blog/pkg/setting"
	"02_go-gin-example/gin-blog/routers/api"
	"02_go-gin-example/gin-blog/routers/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	gin.SetMode(setting.ServerSetting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)     //新增一条记录
		apiv1.PUT("/tags/:id", v1.EditTag) //更新记录
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		apiv1.GET("/articles", v1.GetArticles)          //获取文章列表
		apiv1.GET("/articles/:id", v1.GetArticle)       //获取指定文章
		apiv1.POST("/articles", v1.AddArticle)          //新增文章
		apiv1.PUT("/articles/:id", v1.EditArticle)      //更新文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle) //删除文章
	}
	return r
}
