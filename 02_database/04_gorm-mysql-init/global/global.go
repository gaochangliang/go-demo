package global

import (
	"database/04/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GLOBAL_CONFIG config.Server
	GLOBAL_DB     *gorm.DB
	GLOBAL_VP     *viper.Viper
)
