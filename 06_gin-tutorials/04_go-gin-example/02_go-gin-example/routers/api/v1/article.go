package v1

import (
	"02_go-gin-example/gin-blog/pkg/app"
	"02_go-gin-example/gin-blog/pkg/e"
	"02_go-gin-example/gin-blog/pkg/setting"
	"02_go-gin-example/gin-blog/pkg/util"
	"02_go-gin-example/gin-blog/service/article_service"
	"02_go-gin-example/gin-blog/service/tag_service"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

//type Article struct {
//	TagID int `form:"tag_id" valid:"Min(1)"`
//	State int `form:"State" valid:"Range(0,1)"`
//}

//根据标签和状态获取
func GetArticles(c *gin.Context) {

	/*

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

	*/

	appG := app.Gin{C: c}
	valid := validation.Validation{}

	state := -1

	if arg := c.PostForm("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("state must be 0 or 1")

	}

	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		valid.Min(tagId, 1, "tag_id").Message("tag_id must larger than 1")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	articleService := article_service.Article{
		TagID:    tagId,
		State:    state,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

	articles, err := articleService.GetAll()
	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}

	total, err := articleService.Count()
	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_COUNT_ARTICLE_FAIL, nil)
		return
	}

	data := make(map[string]interface{})

	data["lists"] = articles
	data["total"] = total
	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

//获取单个文章

func GetArticle(c *gin.Context) {
	appG := app.Gin{C: c}
	//注意这里转换的方式
	id := com.StrTo(c.Param("id")).MustInt()

	var valid validation.Validation

	valid.Min(id, 0, "id").Message("id muse large than 0")

	if !valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	articleService := article_service.Article{ID: id}
	exists, err := articleService.ExistByID()

	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	article, err := articleService.Get()
	appG.Response(http.StatusOK, e.SUCCESS, article)
}

//type Article struct {
//	TagID int `form:"tag_id" valid:"Min(1)"`
//	State int `form:"State" valid:"Range(0,1)"`
//}

type AddArticleForm struct {
	TagId         int    `form:"tag_id" valid:"Required,Min(1)"`
	Title         string `form:"title" valid:"Required,MaxSize(100)"`
	Desc          string `form:"desc" valid:"Required,MaxSize(255)"`
	Content       string `form:"content" valid:"Required,MaxSize(65535)"`
	CreateBy      string `form:"create_by" valid:"Required,MaxSize(100)"`
	CoverImageUrl string `form:"cover_image_url" valid:"Required,MaxSize(255)"`
	State         int    `form:"state" valid:"Required,Range(0,1)"`
}

func AddArticle(c *gin.Context) {
	//tagId := com.StrTo(c.Query("tag_id")).MustInt()
	//title := c.Query("title")
	//desc := c.Query("desc")
	//content := c.Query("content")
	//createBy := c.Query("create_by")
	//state := com.StrTo(c.Query("state")).MustInt()
	//
	//valid := validation.Validation{}
	//
	//valid.Min(tagId, 1, "tag_id").Message("tag_id must larger than 1")
	//valid.Required(title, "title").Message("标题不能为空")
	//valid.Required(title, "desc").Message("描述不能为空")
	//valid.Required(title, "content").Message("内容不能为空")
	//valid.Required(title, "createBy").Message("创建人不能为空")
	//valid.Range(state, 0, 1, "state").Message("state 0 or 1")
	//
	//code := e.INVALID_PARAMS
	//if !valid.HasErrors() {
	//	//首先判断标签id是否存在
	//	if models.ExistTagById(tagId) {
	//		data := make(map[string]interface{})
	//		data["tag_id"] = tagId
	//		data["title"] = title
	//		data["desc"] = desc
	//		data["content"] = content
	//		data["created_by"] = createBy
	//		data["state"] = state
	//		models.AddArticle(data)
	//		code = e.SUCCESS
	//	} else {
	//		code = e.ERROR_NOT_EXIST_TAG
	//	}
	//} else {
	//	for _, err := range valid.Errors {
	//		log.Printf("key %v value %v", err.Key, err.Message)
	//	}
	//}
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"code":    code,
	//	"data":    make(map[string]interface{}),
	//	"message": e.GetMsg(code),
	//})

	var (
		appG = app.Gin{C: c}
		form AddArticleForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
	}

	//判断标签是否存在
	tagService := tag_service.Tag{
		ID: form.TagId,
	}
	exist, err := tagService.ExistById()
	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exist {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	//这里不需要获取pageSize, pageNum,不需要返回文章及分页信息
	articleService := article_service.Article{
		TagID:         form.TagId,
		Title:         form.Title,
		Desc:          form.Desc,
		Content:       form.Content,
		CoverImageUrl: form.CoverImageUrl,
		State:         form.State,
		CreateBy:      form.CreateBy,
	}

	if err := articleService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

type EditArticleForm struct {
	ID            int    `form:"id" valid:"Required;Min(1)"`
	TagId         int    `form:"tag_id" valid:"Required,Min(1)"`
	Title         string `form:"title" valid:"Required,MaxSize(100)"`
	Desc          string `form:"desc" valid:"Required,MaxSize(255)"`
	Content       string `form:"content" valid:"Required,MaxSize(65535)"`
	ModifiedBy    string `form:"modified_by" valid:"Required;MaxSize(100)"`
	CoverImageUrl string `form:"cover_image_url" valid:"Required,MaxSize(255)"`
	State         int    `form:"state" valid:"Range(0,1)"`
}

//传递过来的默认是字符串，字符串类型变量不需要处理
//number类型必须要转化

func EditArticle(c *gin.Context) {
	//id := com.StrTo(c.Param("id")).MustInt()
	//tagId := com.StrTo(c.Query("tag_id")).MustInt()
	//title := c.Query("title")
	//desc := c.Query("desc")
	//content := c.Query("content")
	//modifiedBy := c.Query("modified_by")
	//state := com.StrTo(c.Query("state")).MustInt()
	//
	//valid := validation.Validation{}
	//valid.Min(id, 1, "id").Message("id最小值为1")
	//valid.Range(state, 0, 1, "state").Message("state必须为0或者1")
	//valid.MaxSize(title, 100, "title").Message("message最大为100")
	//valid.MaxSize(desc, 255, "desc").Message("desc最大为100")
	//valid.MaxSize(content, 65535, "content").Message("desc最大为65535")
	//valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	//valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	//
	//code := e.INVALID_PARAMS
	//if !valid.HasErrors() {
	//	if models.ExistArticleById(id) {
	//		//需要验证标签
	//		if models.ExistTagById(tagId) {
	//			data := make(map[string]interface{})
	//			if tagId > 0 {
	//				data["tag_id"] = tagId
	//			}
	//			if content != "" {
	//				data["content"] = content
	//			}
	//			if desc != "" {
	//				data["desc"] = desc
	//			}
	//			if title != "" {
	//				data["title"] = title
	//			}
	//			data["modifiedBy"] = modifiedBy
	//			code = e.SUCCESS
	//			models.EditArticle(id, data)
	//		} else {
	//			code = e.ERROR_NOT_EXIST_TAG
	//		}
	//	} else {
	//		code = e.ERROR_NOT_EXIST_ARTICLE
	//	}
	//} else {
	//	for _, err := range valid.Errors {
	//		log.Printf("err.key: %s,err.message:%s", err.Key, err.Message)
	//	}
	//}
	//c.JSON(http.StatusOK, gin.H{
	//	"code":    code,
	//	"data":    make(map[string]string),
	//	"message": e.GetMsg(code),
	//})

	var (
		appG = app.Gin{C: c}
		form = EditArticleForm{
			ID: com.StrTo(c.Param("id")).MustInt(),
		}
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	articleService := article_service.Article{
		ID:            form.ID,
		TagID:         form.TagId,
		Title:         form.Title,
		Desc:          form.Desc,
		Content:       form.Content,
		CoverImageUrl: form.CoverImageUrl,
		State:         form.State,
		ModifiedBy:    form.ModifiedBy,
	}

	exists, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	tagService := tag_service.Tag{ID: form.TagId}
	exists, err = tagService.ExistById()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	err = articleService.Edit()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_EDIT_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

func DeleteArticle(c *gin.Context) {
	//id := com.StrTo(c.Param("id")).MustInt()
	//valid := validation.Validation{}
	//
	//valid.Min(id, 1, "id").Message("最小值为1")
	//code := e.INVALID_PARAMS
	//if !valid.HasErrors() {
	//	if !models.ExistArticleById(id) {
	//		code = e.ERROR_NOT_EXIST_ARTICLE
	//	} else {
	//		code = e.SUCCESS
	//		models.DeleteArticle(id)
	//	}
	//} else {
	//	for _, err := range valid.Errors {
	//		log.Printf("err.key: %s,err.message:%s", err.Key, err.Message)
	//	}
	//}
	//c.JSON(http.StatusOK, gin.H{
	//	"code":    code,
	//	"data":    make(map[string]string),
	//	"message": e.GetMsg(code),
	//})
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("最小值为1")
	if !valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
	}

	articleService := article_service.Article{
		ID: id,
	}
	exists, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	err = articleService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_ARTICLE_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
