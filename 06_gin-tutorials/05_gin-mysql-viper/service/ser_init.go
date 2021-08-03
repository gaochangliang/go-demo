package service

import (
	"database/sql"
	"fmt"
	"gin/05/config"
	"gin/05/global"
	"gin/05/model"
	"gin/05/source"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(conf model.InitDB) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", conf.UserName, conf.Password, conf.Host, conf.Port)
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", conf.DBName)
	if err := createTable(dsn, "mysql", createSql); err != nil {
		return err
	}

	//gorm建立链接
	MysqlConfig := config.Mysql{
		Path:     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		Dbname:   conf.DBName,
		Username: conf.UserName,
		Password: conf.Password,
		Config:   "charset=utf8mb4&parseTime=True&loc=Local",
	}

	linkDns := MysqlConfig.Username + ":" + MysqlConfig.Password + "@tcp(" + MysqlConfig.Path + ")/" + MysqlConfig.Dbname + "?" + MysqlConfig.Config
	mysqlConfig := mysql.Config{
		DSN:                       linkDns, // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(MysqlConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(MysqlConfig.MaxOpenConns)
		global.GLOBAL_DB = db
	}

	err := global.GLOBAL_DB.AutoMigrate(
		model.SysUser{},
	)
	if err != nil {
		global.GLOBAL_DB = nil
		return err
	}

	//初始化表数据，添加超级管理员
	err = initTable(
		source.Admin,
	)
	if err != nil {
		global.GLOBAL_DB = nil
		return err
	}
	return nil
}

func createTable(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		var err = db.Close()
		if err != nil {
			return
		}
	}(db)

	if err := db.Ping(); err != nil {
		return err
	}

	_, err = db.Exec(createSql)
	return err
}

func initTable(initTableFunctions ...model.InitTableFunc) error {
	for _, v := range initTableFunctions {
		err := v.Init()
		if err != nil {
			return err
		}
	}
	return nil
}
