// Package global Some configurations for primary storage reads
package global

import (
	"gin/05/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GLOBAL_DB     *gorm.DB
	GLOBAL_VIPER  *viper.Viper
	GLOBAL_CONFIG *config.Server
)
