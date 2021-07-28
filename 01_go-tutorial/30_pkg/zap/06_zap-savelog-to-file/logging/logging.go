package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var Global *zap.SugaredLogger

var (
	logLevel = zapcore.InfoLevel
)

func init() {

	writeSyncer := getLogWriter() //Where the logs are written to
	encoder := getEncoder()       //encoder:How to write logs

	core := zapcore.NewCore(encoder, writeSyncer, logLevel)
	logger := zap.New(core)
	logger = logger.WithOptions(zap.AddCaller())
	Global = logger.Sugar()

}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig() //Creating a default encoder configuration
	encoderConfig.NameKey = "name"
	//encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder //modify time encoder
	encoderConfig.EncodeTime = customTimeFormat //custom time encoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func customTimeFormat(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 - 15:04:05"))
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)
}
