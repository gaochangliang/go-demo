package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	ExportSavePath  string
	QrCodeSavePath  string
	FontSavePath    string
	RuntimeRootPath string
	PrefixUrl       string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

func SetUp() {
	//load file
	Cfg, err := ini.Load("conf/app.ini") //配置文件首字母大写才能读到
	//Cfg, err := ini.InsensitiveLoad("conf/app.ini")	//不区配置文件子选项大小写
	if err != nil {
		log.Printf("failed to load 'conf/app.ini'  err %v", err)

	}

	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Printf("cfg section app mapto AppSetting err %v", err)
	}

	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Printf("cfg section server mapto ServerSetting err %v", err)
	}

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

}
