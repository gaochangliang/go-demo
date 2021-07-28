package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"zap/06/config"
)

var (
	GVIPER *viper.Viper
	CONFIG config.Server
	ZAPLOG *zap.Logger
)
