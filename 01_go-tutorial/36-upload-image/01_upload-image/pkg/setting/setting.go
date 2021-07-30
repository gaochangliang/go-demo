package setting

import (
	"fmt"
	"github.com/go-ini/ini"
	"log"
)

type app struct {
	PageSize        int
	jwtSecret       int
	RuntimeRootPath string
	ImagePrefixUrl  string
	ImageSavePath   string
	ImageMaxSize    int
	ImageAllowExts  []string
	LogSavePath     string
	LogSaveName     string
	LogFileExt      string
	TimeFormat      string
}

var AppSetting = &app{}

func SetUp() {
	Cfg, err := ini.Load("config/app.ini")
	if err != nil {
		log.Printf("fail to load config/app.ini err %v", err)
	}

	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Printf("cfg section error %v", err)
	}
	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	fmt.Println(AppSetting)
}
