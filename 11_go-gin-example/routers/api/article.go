package v1

import (
	"gin-blog/models"
	"gin-blog/pkg/e"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

//根据标签和状态获取
func GetArticles(c *gin.Context) {

	//返回的数据
	data := make(map[string]interface{})
	//查询条件
	maps := make(map[string]interface{})

	valid := validation.Validation{}

	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		valid.Min(tagId, 1, "tag_id").Message("tag_id must larger than 1")
		maps["tag_Id"] = tagId
	}

	var state = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("state must be 0 or 1")
		maps["state"] = state
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		data["list"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
	} else {
		for _, err := range valid.Errors {
			log.Printf("key %v err %v", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,           //数据码
		"data":    data,           //具体返回的数据
		"message": e.GetMsg(code), //消息
	})
}

//获取单个文章

func GetArticle(c *gin.Context) {
	var data interface{}
	//注意这里转换的方式
	id := com.StrTo(c.Param("id")).MustInt()

	var valid validation.Validation

	valid.Min(id, 0, "id").Message("id muse large than 0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		//需要判断id是否存在
		if !models.ExistArticleById(id) {
			code = e.ERROR_NOT_EXIST_ARTICLE
		} else {
			code = e.SUCCESS
			data = models.GetArticle(id)
		}

	} else {
		for _, err := range valid.Errors {
			log.Printf("key %v err %v", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    data,
		"message": e.GetMsg(code),
	})
}

func addArticle(c *gin.Context) {
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createBy := c.Query("create_by")
	state := com.StrTo(c.Query("state")).MustInt()

	valid := validation.Validation{}

	valid.Min(tagId, 1, "tag_id").Message("tag_id must larger than 1")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(title, "desc").Message("描述不能为空")
	valid.Required(title, "content").Message("内容不能为空")
	valid.Required(title, "createBy").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("state 0 or 1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		//首先判断标签id是否存在
		if models.ExistTagById(tagId) {
			data := make(map[string]interface{})
			data["tag_id"] = tagId
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["create_by"] = createBy
			data["state"] = state
			models.AddArticle(data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("key %v value %v", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    make(map[string]interface{}),
		"message": e.GetMsg(code),
	})

}



func EditArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createBy := c.Query("create_by")
	state := com.StrTo(c.Query("state")).MustInt()

	valid := validation.Validation{}

	valid.


}
