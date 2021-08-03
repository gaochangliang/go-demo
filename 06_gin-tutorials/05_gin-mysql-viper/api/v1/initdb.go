package v1

import (
	"gin/05/app"
	"gin/05/e"
	"gin/05/global"
	"gin/05/model"
	"gin/05/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitDB(c *gin.Context) {
	var g = &app.Gin{C: c}
	//已经可以连接数据库
	if global.GLOBAL_DB != nil {
		g.Response(http.StatusInternalServerError, e.ErrorDBExist, nil)
		return
	}

	var dbInfo model.InitDB
	if err := c.ShouldBindJSON(&dbInfo); err != nil {
		g.Response(http.StatusInternalServerError, e.InvalidParams, nil)
		return
	}

	if err := service.InitDB(dbInfo); err != nil {
		g.Response(http.StatusInternalServerError, e.InvalidParams, nil)
		return
	}
	g.Response(http.StatusOK, e.Success, nil)
}
