package main

import (
	"encoding/json"
	"go.uber.org/zap"
)

func main() {
	rawJSON := []byte(`{
    "level":"debug",
    "encoding":"json",
    "outputPaths": ["stdout", "server.log"],
    "errorOutputPaths": ["stderr"],
    "initialFields":{"name":"dj"},
    "encoderConfig": {
      "messageKey": "message",
      "levelKey": "level",
      "levelEncoder": "lowercase"
    }
  }`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}

	//cfg的配置中
	//Level：日志级别；
	//Encoding：输出的日志格式，默认为 JSON；
	//OutputPaths：可以配置多个输出路径，路径可以是文件路径和stdout（标准输出）；
	//ErrorOutputPaths：错误输出路径，也可以是多个；
	//InitialFields：每条日志中都会输出这些值。
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	logger.Info("server start work successfully")
	logger.Info("server start work successfully")
}
