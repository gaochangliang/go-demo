package v1

import (
	"fmt"
	"gin-blog/models"
	"gin-blog/pkg/e"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

// @Summary 测试获取标签
// @Tags 测试
// @Description 获取文章标签
// @Accept json
// @Param   name  query string true     "人名"
// @Success 200 {string} string "{"msg": "hello Razeen"}"
// @Failure 400 {string} string "{"msg": "who are you"}"
// @Router /api/v1/GetTags [get]
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

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//新增文件标签
func AddTag(c *gin.Context) {
	name := c.Query("name")

	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()

	fmt.Println("state", state)
	createdBy := c.Query("create_by")
	fmt.Println("create_By", createdBy)

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最大为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许为0或1")

	//Invalid parameters are always returned if there is an error

	code := e.INVALID_PARAMS
	fmt.Println("*********")
	if !valid.HasErrors() {
		fmt.Println("1")
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

//修改文章标签
func EditTag(c *gin.Context) {
	//获取传入的参数
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("state 0/1")
	}

	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.Required(id, "id").Message("id不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("name最长为100")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if !models.ExistTagById(id) {
			code = e.ERROR_NOT_EXIST_TAG
		} else {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			data["state"] = state
			models.EditTag(id, data)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

//删除文章标签
func DeleteTag(c *gin.Context) {
	// /tags/:id 这种形式路由必须有param形式处理
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	//id必须大于0
	valid.Min(id, 1, "id").Message("id must larger than 0")

	code := e.INVALID_PARAMS

	//看出什么错误
	//for _, v := range valid.Errors {
	//	fmt.Println("v", v)
	//}

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagById(id) {
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string), //虽然这里不返回具体的数据，但是保证返回的字段一样，所以加上
	})
}
