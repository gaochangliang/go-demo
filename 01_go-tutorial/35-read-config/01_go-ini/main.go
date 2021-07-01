package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"time"
)

/*

#debug or release
RUN_MODE = debug

[app]
PAGE_SIZE = 10
JWT_SECRET = 23347$040412

[server]
HTTP_PORT = 8000
READ_TIMEOUT = 60
WRITE_TIMEOUT = 60

[database]
TYPE = mysql
USER = 数据库账号
PASSWORD = 数据库密码
#127.0.0.1:3306
HOST = 数据库IP:数据库端口号
NAME = blog
TABLE_PREFIX = blog_

*/

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

var (
	dbType, dbName, user, password, host, tablePrefix string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")

	//如果需要不区分键值大小写的话
	//cfg, err := ini.InsensitiveLoad("conf/app.ini")

	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	//加载基本的变量
	LoadBase()
	//加载server区域
	LoadServer()
	//加载app配置
	LoadApp()
	//加载数据库
	LoadDataBase()
}

func main() {
	_ = Cfg.SaveTo("my.ini.local")
}

func LoadBase() {
	//默认分区section用空值读取
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	//为空的话默认设置值为8000，同时进行string->int类型转换
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

func LoadDataBase() {
	sec, err := Cfg.GetSection("database")

	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	//我们可以做一些候选值限制的操作
	//如果读取的值不在候选列表内，则会回退使用提供的默认值,比如下面的类型要求只能用redis或者mysql，都不是默认用mysql
	dbType = sec.Key("TYPE").In("mysql", []string{"mysql", "redis"})

	//型读取操作，默认分区可以使用空字符串表示
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()

	//同时可以对读取的值进行验证
	tablePrefix = sec.Key("TABLE_PREFIX").Validate(func(in string) string {
		fmt.Println("TABLE_PREFIX", in)
		if len(in) > 2 {
			return "my"
		}
		return in
	})

}

//文档  https://ini.unknwon.cn/docs/intro/getting_started
