package main

import "go.uber.org/zap"

func main() {
	logger, _ := zap.NewProduction(zap.AddCaller())
	defer logger.Sync()
	logger.Info("hello world")

	//zap.AddCallerSkip(skip int) 打印行数向上追溯一层
	logger2, _ := zap.NewProduction(zap.AddCaller(), zap.AddCallerSkip(1))
	defer logger2.Sync()
	//替换zap报自定义的全局变量
	zap.ReplaceGlobals(logger2)
	Output("hello world")
}

func Output(msg string, fields ...zap.Field) {
	zap.L().Info(msg, fields...)
}
