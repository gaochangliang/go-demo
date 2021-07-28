//Customize logs, modify log configuration such as default time format,
//logs stored to file
package main

import "zap/06/logging"

func main() {
	logging.Global.Infof("save log to file")
}

//default message
//{"level":"info","ts":1627463529.6973472,"caller":"06_zap-savelog-to-file/main.go:6","msg":"save log to file"}

//need to define
