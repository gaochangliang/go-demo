package main

import (
	"example/uploadimage/pkg/setting"
	gin "github.com/gin-gonic/gin"
)

func main() {
	setting.SetUp()
	r := gin.Default()

	r.POST("/upload", api.UploadImage)

}
