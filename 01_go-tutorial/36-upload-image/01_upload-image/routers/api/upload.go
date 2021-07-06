package api

import (
	"example/uploadimage/pkg/e"
	"example/uploadimage/pkg/upload"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UploadImage(c *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]string)

	//gin封装上传文件方式
	//file文件的指针，
	//image包含了一些image的信息 image.Filename可以获取文件名字
	file, image, err := c.Request.FormFile("image")

	//log.Println("c.Request.FormFile file", file)
	//log.Println("c.Request.FormFile image", image)

	if err != nil {
		log.Printf("c.Request.FormFile err: %s\n", err)
		code = e.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}

	if image == nil {
		code = e.InvalidParams
	} else {
		imageName := upload.GetImageName(image.Filename) // 6998bbc0e3e0727415e140016552bc86.png
		fullPath := upload.GetImageFullPath()            //	runtime/upload/images/
		savePath := upload.GetImagePath()                // upload/images/

		src := fullPath + imageName // runtime/upload/images/6998bbc0e3e0727415e140016552bc86.png
		//判断文件大小和后缀
		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			code = e.Error_UPLOAD_CHECK_IMAGE_FORMAT
		} else {
			//创建目录以及判断是否对目录有操作权限
			err := upload.CheckImage(fullPath)
			if err != nil {
				log.Println("CheckImage err", err)
			} else if err := c.SaveUploadedFile(image, src); err != nil {
				log.Println("c.SaveUploadedFile err", err)
			} else {
				data["image_url"] = upload.GetImageFullUrl(imageName)
				data["image_save_url"] = savePath + imageName
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
