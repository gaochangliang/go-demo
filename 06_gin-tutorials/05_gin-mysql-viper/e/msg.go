package e

var msgFlags = map[int]string{
	Success:                 "ok",
	Error:                   "fail",
	InvalidParams:           "请求参数错误",
	ErrorAddUserFail:        "添加用户失败",
	ErrorCheckExistUserFail: "检查用户失败",

	ErrorDBExist: "已存在数据库配置",
}

func GetMsg(code int) string {
	msg, ok := msgFlags[code]
	if !ok {
		return msgFlags[Error]
	}
	return msg
}
