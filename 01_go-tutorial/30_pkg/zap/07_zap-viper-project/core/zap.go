package core

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
	"zap/07/global"
	"zap/07/utils"
)

var level zapcore.Level

func Zap() (logger *zap.Logger) {
	if ok, err := utils.CreateDir(global.CONFIG.Zap.Director); !ok {
		fmt.Println("CreateDir Err:", err)
	}

	switch global.CONFIG.Zap.Level { // 初始化配置文件的Level
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	//We will use the zap.New(...) method to pass all the configuration manually,
	//instead of using a preconfigured method like zap.NewProduction() to create the logger
	if level == zap.DebugLevel || level == zap.ErrorLevel {
		//The first parameter of the new method is zapcore.Core,
		//and the zapcore.NewCore method shows that it takes three parameters

		//Encoder: Encoder (how to write logs)
		//WriterSyncer : Specify where the logs will be written to
		//Log Level: Which level of logs will be written
		logger = zap.New(getEncoderCore(), zap.AddStacktrace(level))
	} else {
		logger = zap.New(getEncoderCore())
	}

	if global.CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	return logger
}

func getEncoderCore() (core zapcore.Core) {
	writer, err := utils.GetWriteSyncer()
	if err != nil {
		fmt.Printf("get write Syncer err:%v", err.Error())
		return
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}

func getEncoder() zapcore.Encoder {
	if global.CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}

	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.CONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case global.CONFIG.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.CONFIG.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.CONFIG.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.CONFIG.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.CONFIG.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}
