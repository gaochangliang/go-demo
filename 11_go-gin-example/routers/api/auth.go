package api

import (
	"gin-blog/models"
	"gin-blog/pkg/e"
	"gin-blog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	var auth = auth{
		Username: username,
		Password: password,
	}

	valid := validation.Validation{}
	ok, _ := valid.Valid(&auth)
	code := e.INVALID_PARAMS
	data := make(map[string]interface{})
	if ok {
		isExist := models.CheckAuth(username, password)
		if !isExist {
			code = e.ERROR_AUTH
		} else {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("key %s val %s\n", err.Key, err.Value)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  e.GetMsg(code),
	})
}
