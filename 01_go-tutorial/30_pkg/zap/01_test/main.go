package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	logger := zap.NewExample()
	defer logger.Sync()

	url := "kobe"
	logger.Info("fail to fetch url",
		//由于fmt.Printf之类的方法大量使用interface{}和反射，会有不少性能损失，并且增加了内存分配的频次。
		//zap为了提高性能、减少内存分配次数，没有使用反射
		//且默认的Logger只支持强类型的、结构化的日志
		//zap为 Go 语言中所有的基本类型和其他常见类型都提供了方法
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	//每个字段都用方法包一层用起来比较繁琐。zap也提供了便捷的方法SugarLogger，
	//可以使用printf格式符的方式。调用logger.Sugar()即可创建SugaredLogger。
	//SugaredLogger的使用比Logger简单，只是性能比Logger低 50% 左右，可以用在非热点函数中。
	//调用SugarLogger以f结尾的方法与fmt.Printf没什么区别，比如下面的Infof
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)

	sugar.Infof("Failed to fetch URL")
}
