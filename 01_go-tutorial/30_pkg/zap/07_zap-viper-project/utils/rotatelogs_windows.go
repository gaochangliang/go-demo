package utils

import (
	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
	"zap/07/global"
)

func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		path.Join(global.CONFIG.Zap.Director, "%Y-%m-%d.log"),
		zaprotatelogs.WithMaxAge(7*time.Minute),
		zaprotatelogs.WithRotationTime(1*time.Minute),
	)

	if global.CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}

	return zapcore.AddSync(fileWriter), err
}
