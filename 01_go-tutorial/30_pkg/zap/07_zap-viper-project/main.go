package main

import (
	"zap/07/core"
	"zap/07/global"
)

func main() {
	global.GVIPER = core.Viper()
	global.ZAPLOG = core.Zap()

	global.ZAPLOG.Info("1")
	global.ZAPLOG.Info("12")
	global.ZAPLOG.Error("123")

}
