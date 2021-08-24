package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

// @Summary 获取文章(介绍)
// @Tags    文章
// @Produce json
// @Param name query string true "Name"
// @Param state query int   false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 400 {string} string "{"msg": "who are you"}"
// @Router /api/v1/tags [get]
func GetTags(c *gin.Context) {

	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "test",
		"data": data,
	})
}


func AddTag(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    "456",
		"message": "AddTag",
		"data":    make(map[string]string),
	})
}