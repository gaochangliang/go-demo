package main

import (
	"fmt"
	"time"
)

//     log20060102.log
var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSavePath, time.Now().Format(TimeFormat), LogFileExt)

	fmt.Println(prefixPath + suffixPath)
}

func main() {

}
