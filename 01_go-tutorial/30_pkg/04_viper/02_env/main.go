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

	//函数里面只有一个参数，变量名为 SPF_ID 绑定的环境变量也为 SPF_ID
	_ = viper1.BindEnv("id")            //
	_ = viper1.BindEnv("GOP", "GOPATH") // 相当于SPF_GOP绑定GOPATH

	_ = os.Setenv("SPF_ID", "13")  // typically done outside of the app
	_ = os.Setenv("GOPATH", "144") // typically done outside of the app

	//对于环境变量自身来说，是区分大小写的
	//BindEnv只有一个参数时，默认gopath既是变量名又是环境变量名称
	//所以下面viper.Get("gopath")获取不到GOPATH的值
	_ = viper1.BindEnv("gopath")

	//在使用Get的时候，viper 会自动加上这个前缀再从环境变量中查找。
	//如果对应的环境变量不存在，viper 会自动将键名全部转为大写再查找一次。
	//所以，使用键名gopath也能读取环境变量GOP的值
	fmt.Println("gop", viper1.Get("id"))     // 13
	fmt.Println("gopath", viper1.Get("gop")) // 13
}

func do2() {
	//viper1和viper2是独立的，没有关系
	fmt.Println("gopath", viper2.Get("gop")) // 13
}

//viper可以访问环境变量的值
//
