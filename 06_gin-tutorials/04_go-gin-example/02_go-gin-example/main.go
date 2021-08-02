package main

import (
	"02_go-gin-example/gin-blog/models"
	"02_go-gin-example/gin-blog/pkg/logging"
	"02_go-gin-example/gin-blog/pkg/setting"
	"02_go-gin-example/gin-blog/pkg/upload"
	"02_go-gin-example/gin-blog/routers"
	"02_go-gin-example/gin-blog/routers/api"
	"fmt"
	"net/http"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample api celler api.
// @termsOfService https://www.topgoer.com
// @contact.name www.topgoer.com
// @contact.url https://www.topgoer.com
// @contact.email me@razeen.me
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:80
// @BasePath /api/v1
func main() {

	setting.Setup()
	models.Setup()
	logging.Setup()

	//加入service tag article 路由只需要处理参数和拿到结果 -> 把所有的逻辑加到服务里面去 -> 数据库层只查询数据
	r := routers.InitRouter()
	r.POST("/upload", api.UploadImage)
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        r,
		ReadTimeout:    0,
		WriteTimeout:   0,
		MaxHeaderBytes: 0,
	}
	_ = s.ListenAndServe()
}
