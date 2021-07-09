package app

import (
	"example/01_qrcode/pkg/e"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}
