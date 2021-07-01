##  swaggo简介

主要用于自动生成文档，支持gin,echo等框架,

```
https://github.com/swaggo/swag
```

## 安装swag

```
go get -u github.com/swaggo/swag/cmd/swag
```

验证是否安装成功

```
swag -v
```

## 安装gin-swagger

```go
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
go get -u github.com/alecthomas/template
```

## 操作流程
Swagger 中需要将相应的注释或注解编写到方法上，再利用生成器自动生成说明文件



**路由**

```
import (
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	...
	
)

func main(){
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
```

**编写注释**

```

// @Summary 测试获取标签
// @Tags 测试
// @Description 获取文章标签
// @Accept json
// @produce json
// @Param   name  query string true     "人名"
// @Success 200 {string} string "{"msg": "hello Razeen"}"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 400 {string} string "{"msg": "who are you"}"
// @Router /api/v1/GetTags [get]
func AddTag(c *gin.Context) {

}

//summary 		路由介绍标题
//Tags	  		路由归类标签
//Description   简介描述
//produce       返回格式json
//param     	name(参数名)  query(参数传递形式) string（参数type） true(是否必传)    "人名"（参数介绍）
//Success成功返回  200状态码   返回的数据类型  返回的数据具体类型  "{"msg": "hello Razeen"}"具体的内容           
//Failure错误返回  400状态码   返回的数据类型是一个对象  返回的数据具体类型  web.APIError  具体的内容           
```

**生成文档**

根目录执行

```
swag init
```

每次`API`注解修改完成之后都需要重新生成

**main.go引用文档**

这里使用了go mod，参考

```
import (
	...
	_ "gin-doc/docs"
	...
)
```



## 例子

```
/ @Summary Get multiple articles
// @Produce  json
// @Param tag_id body int false "TagID"
// @Param state body int false "State"
// @Param created_by body int false "CreatedBy"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/articles [get]
func GetArticles(c *gin.Context) {
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, data)
}
```

appG.Response调用如下方法,Response为就是app.Response

```
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}
```

app.Response是一个结构体

```
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
```

