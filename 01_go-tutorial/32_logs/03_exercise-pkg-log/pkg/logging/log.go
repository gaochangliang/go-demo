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
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""                                                  //prefix定义每个生成的日志行的开头
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"} //flag定义了日志记录属性
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)

	//创建一个新的日志记录器
	//out定义要写入日志数据的IO句柄。
	//prefix定义每个生成的日志行的开头。
	//flag定义了日志记录属性
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
	setPrefix(WARNING)
	logger.Println(v)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}

func setPrefix(level Level) {
	//DefaultCallerDepth 错误的发生往上追溯的层数
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		//Base函数返回路径的最后一个元素  "/abc/" => abc  /abc->abc
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
