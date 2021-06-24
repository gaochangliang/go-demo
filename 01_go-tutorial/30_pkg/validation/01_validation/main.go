package main

import (
	"github.com/astaxie/beego/validation"
	"log"
)

func main() {
	name := "zhaowei"
	name2 := ""
	valid := validation.Validation{}

	//name is the array to be verified, key is the string to be prompted
	valid.Required(name, "name").Message("name不能为空")
	valid.Required(name2, "name3").Message("name2不能为空")

	valid.MaxSize(name, 2, "namemax").Message("name的最大长度不能超过2")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	if v := valid.Max(name, 3, "name3"); !v.Ok {
		log.Println("")
	}

}

//Mainly used to verify that the data meets the requirements
