package initalize

import (
	"database/04/global"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.GLOBAL_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}
