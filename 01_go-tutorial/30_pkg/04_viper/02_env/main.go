package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var (
	//viper1 和 viper2 是互不相关的
	viper1 *viper.Viper
	viper2 *viper.Viper
)

func init() {
	// 绑定环境变量
	viper1 = viper.New()
	viper2 = viper.New()

	viper1.AutomaticEnv()
	viper2.AutomaticEnv()

}

func main() {
	do1()
	do2()
}

func do1() {
	viper1.SetEnvPrefix("spf") // will be uppercased automatically
	_ = viper1.BindEnv("id")   //因为这个id没有绑定spf前缀

	_ = os.Setenv("SPF_ID", "13") // typically done outside of the app

	//对于环境变量自身来说，是区分大小写的
	//BindEnv只有一个参数时，默认gopath既是变量名又是环境变量名称
	//所以下面viper.Get("gopath")获取不到GOPATH的值
	_ = viper1.BindEnv("gopath")
	fmt.Println("gopath", viper1.Get("id"))     // 13
	fmt.Println("gopath", viper1.Get("gopath")) // 13
}

func do2() {

}

//viper可以访问环境变量的值
//
