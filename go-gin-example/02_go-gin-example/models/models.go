package models

import (
	"02_go-gin-example/gin-blog/pkg/setting"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`  //创建时间
	ModifiedOn int `json:"modified_on"` //修改时间
	DeletedOn  int `json:"deleted_on"`  //删除
}

func Setup() {

	db, err := gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", setting.DatabaseSetting.User, setting.DatabaseSetting.Password, setting.DatabaseSetting.Host, setting.DatabaseSetting.Name))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

}

func CloseDB() {
	defer db.Close()
}
