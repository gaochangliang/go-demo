package api

import (
	"example/uploadimage/pkg/e"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func uploadImage(c *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]string)

	file, image, err := c.Request.FormFile("image")
	if err != nil {
		log.Printf("c.Request.FormFile err: %s\n", err)
		code = e.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}

}
