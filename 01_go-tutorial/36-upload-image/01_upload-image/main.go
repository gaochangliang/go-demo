package main

import (
	"example/uploadimage/pkg/setting"
	"example/uploadimage/pkg/upload"
	"example/uploadimage/routers/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	setting.SetUp()
	r := gin.Default()

	r.POST("/upload", api.UploadImage)
	//添加处理图片的逻辑
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	_ = r.Run(":8000")

}
