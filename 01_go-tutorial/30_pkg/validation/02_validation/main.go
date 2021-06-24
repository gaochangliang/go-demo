package main

import (
	"github.com/astaxie/beego/validation"
	"log"
)

func main() {
	name := "zhaowei"
	name2 := ""
	valid := validation.Validation{}
	valid.Required(name, "name").Message("name不能为空")
	valid.Required(name2, "name3").Message("name2不能为空")

	valid.MaxSize(name, 2, "namemax").Message("name的最大长度不能超过2")

	if v := valid.Max(name, 3, "name3"); !v.Ok {
		log.Println(v.Error.Key, v.Error.Message)
	}

}
