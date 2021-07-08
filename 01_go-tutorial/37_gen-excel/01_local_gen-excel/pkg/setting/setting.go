package setting

import (
	"fmt"
	"github.com/go-ini/ini"
	"log"
)

type Excel struct {
	RuntimeRootPath string
	ExportSavePath  string
	PrefixUrl       string
}

var ExcelSetting = &Excel{}

func Setup() {
	Cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	err = Cfg.Section("excel").MapTo(ExcelSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo ExcelSetting err: %v", err)
	}

	fmt.Println("ExcelSetting", ExcelSetting)
}
