// Package initalize 主要用于做路由初始化，数据库表初始化等工作
package initalize

import (
	"gin/05/model"
	"gorm.io/gorm"
	"log"
	"os"
)

func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.SysUser{},
	)
	if err != nil {
		log.Println("MysqlTables error", err)
		os.Exit(0)
	}

	log.Println("register table success")
}

