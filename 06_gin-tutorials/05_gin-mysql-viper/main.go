//主要结合viper gin mysql 自定义err 做一个数据库的CURD操作
package main

import (
	"gin/05/core"
	"gin/05/global"
	"gin/05/initalize"
)

func main() {
	global.GLOBAL_VIPER = core.Viper()  // 初始化Viper
	global.GLOBAL_DB = core.GormMysql() // gorm连接数据库
	if global.GLOBAL_DB != nil {
		initalize.MysqlTables(global.GLOBAL_DB) // 初始化表
		db, _ := global.GLOBAL_DB.DB()
		defer db.Close()
	}
	core.RunServer()

}


