package model

type InitDB struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
	DBName   string `json:"dbName"`
}

//后面form字段参考  https://github.com/go-playground/form
//将url.Values解码为Go值，将Go值编码为url.Values。

//相关的验证器字段列表
//https://pkg.go.dev/github.com/go-playground/validator?utm_source=godoc#hdr-Baked_In_Validators_and_Tags
