package main

import "go.uber.org/zap"

func main() {
	//zap提供了两个全局的Logger，一个是*zap.Logger， 可调用zap.L()获得；
	//另一个是*zap.SugaredLogger，可调用zap.S()获得
	//这两个zap包的Logger默认并不会记录日志！它是一个无实际效果的Logger
	zap.L().Info("global logger before")
	zap.S().Info("global SugaredLogger before")

	logger := zap.NewExample()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	zap.L().Info("global logger after")
	zap.S().Error("global SugaredLogger after")
}
