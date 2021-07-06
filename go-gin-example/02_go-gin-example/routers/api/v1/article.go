package v1

import (
	"02_go-gin-example/gin-blog/models"
	"02_go-gin-example/gin-blog/pkg/app"
	"02_go-gin-example/gin-blog/pkg/e"
	"02_go-gin-example/gin-blog/pkg/logging"
	"02_go-gin-example/gin-blog/pkg/setting"
	"02_go-gin-example/gin-blog/pkg/util"
	"02_go-gin-example/gin-blog/service/article_service"
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
		data["list"] = models.GetArticles(util.GetPage(c), setting.AppSetting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
	} else {
		for _, err := range valid.Errors {
			logging.Info("key", err.Key, "message", err.Message)
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
	appG := app.Gin{C: c}
	var data interface{}
	//注意这里转换的方式
	id := com.StrTo(c.Param("id")).MustInt()

	var valid validation.Validation

	valid.Min(id, 0, "id").Message("id muse large than 0")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
	}

	articleService := article_service.Article{ID: id}
	exists, err := articleService.ExistByID()

	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
	}

	article, err := articleService.Get(id)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    data,
		"message": e.GetMsg(code),
	})
}

func AddArticle(c *gin.Context) {
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
			data["created_by"] = createBy
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

//传递过来的默认是字符串，字符串类型变量不需要处理
//number类型必须要转化

func EditArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")
	state := com.StrTo(c.Query("state")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id最小值为1")
	valid.Range(state, 0, 1, "state").Message("state必须为0或者1")
	valid.MaxSize(title, 100, "title").Message("message最大为100")
	valid.MaxSize(desc, 255, "desc").Message("desc最大为100")
	valid.MaxSize(content, 65535, "content").Message("desc最大为65535")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleById(id) {
			//需要验证标签
			if models.ExistTagById(tagId) {
				data := make(map[string]interface{})
				if tagId > 0 {
					data["tag_id"] = tagId
				}
				if content != "" {
					data["content"] = content
				}
				if desc != "" {
					data["desc"] = desc
				}
				if title != "" {
					data["title"] = title
				}
				data["modifiedBy"] = modifiedBy
				code = e.SUCCESS
				models.EditArticle(id, data)
			} else {
				code = e.ERROR_NOT_EXIST_TAG
			}
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s,err.message:%s", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    make(map[string]string),
		"message": e.GetMsg(code),
	})

}

func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}

	valid.Min(id, 1, "id").Message("最小值为1")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistArticleById(id) {
			code = e.ERROR_NOT_EXIST_ARTICLE
		} else {
			code = e.SUCCESS
			models.DeleteArticle(id)
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s,err.message:%s", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    make(map[string]string),
		"message": e.GetMsg(code),
	})
}
