//create initialize gorm mysql
package main

import (
	"database/04/core"
	"database/04/global"
	"database/04/initalize"
	"time"
)

func main() {
	core.Viper()
	global.GLOBAL_DB = initalize.Gorm()

	if global.GLOBAL_DB != nil {
		initalize.MysqlTables(global.GLOBAL_DB)
		db, _ := global.GLOBAL_DB.DB()
		//程序结束前关闭链接
		defer db.Close()
	}

	time.Sleep(10 * time.Second)
}
