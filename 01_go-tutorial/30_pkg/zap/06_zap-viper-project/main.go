package main

import (
	"zap/06/core"
	"zap/06/global"
)

func main() {
	global.GVIPER = core.Viper()
	global.ZAPLOG = core.Zap()

	global.ZAPLOG.Info("1")
	global.ZAPLOG.Info("12")
	global.ZAPLOG.Error("123")

}
