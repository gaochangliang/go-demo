package core

import (
	"gin/05/global"
	"github.com/spf13/viper"
	"log"
)

func Viper(path ...string) *viper.Viper {
	v := viper.New()
	v.SetConfigName("config") // name of config file (without extension)
	v.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name

	//also we can set filename
	//viper.SetConfigFile("./config/config.yaml")

	//search paths can be set more than one, viper will look for them in order according to the set order.
	v.AddConfigPath("../config") // call multiple times to add many search paths

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found; ignore error if desired", err)
		} else {
			log.Fatal("Config file was found but another error was produced", err)
		}
	}
	if err := v.Unmarshal(&global.GLOBAL_CONFIG); err != nil {
		log.Fatalf("Viper v.Unmarshal err:%v", err)
	}
	return v
}
