package setting

import (
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string //Note the picture format here

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DataBaseSetting = &Database{}

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

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Printf("cfg section server mapto ServerSetting err %v", err)
	}

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	err = Cfg.Section("database").MapTo(DataBaseSetting)
	if err != nil {
		log.Printf("cfg section database mapto databaseSetting err %v", err)
	}

	fmt.Println("serverSetting", ServerSetting)
	fmt.Println("dataBaseSetting", DataBaseSetting)
	fmt.Println("appSetting", AppSetting)

}
