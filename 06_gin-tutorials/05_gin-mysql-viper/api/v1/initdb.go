package v1

import (
	"gin/05/app"
	"gin/05/e"
	"gin/05/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitDB(c *gin.Context) {
	var g = &app.Gin{C: c}
	if global.GLOBAL_DB != nil {
		g.Response(http.StatusInternalServerError, e.ErrorDBExist, nil)
		return
	}
}
