package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	logger *log.Logger
	F      *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logPrefix  = "" //			prefix定义每个生成的日志行的开头。flag定义了日志记录属性
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARING
	ERROR
	FATAL
)

func init() {
	filePath := getLogFileFullPath()
	F := openLogFile(filePath)

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

func Warn(v ...interface{}) {
	setPrefix(WARING)
	logger.Println(v)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Println(v)
}

//[INFO][article.go:48]2021/06/29 10:53:11 [key state message state must be 0 or 1]
func setPrefix(level Level) {
	//获取代码运行的文件行数
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
