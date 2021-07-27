package main

import "go.uber.org/zap"

func main() {

	//如果每条日志都要记录一些共用的字段，那么使用zap.Fields(fs ...Field)创建的选项
	logger := zap.NewExample(zap.Fields(
		zap.Int("serverId", 41),
		zap.String("player", "kobe"),
	))

	logger.Info("hello world")
}
