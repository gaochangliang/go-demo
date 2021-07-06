package e

var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	InvalidParams:                   "请求参数错误",
	ErrorNotExistTag:                "该标签不存在",
	ErrorAuthCheckTokenFail:         "Token鉴权失败",
	ErrorAuthCheckTokenTimeout:      "Token已超时",
	Error_UPLOAD_CHECK_IMAGE_FAIL:   "检查图片失败",
	Error_UPLOAD_CHECK_IMAGE_FORMAT: "图片格式或者大小有问题",
	Error_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
