package core

import (
	"database/04/global"
	"fmt"
	"github.com/spf13/viper"
)

func Viper(path ...string) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name

	//also we can set filename
	//viper.SetConfigFile("./config/config.yaml")

	//search paths can be set more than one, viper will look for them in order according to the set order.
	viper.AddConfigPath(".") // call multiple times to add many search paths

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found; ignore error if desired", err)
		} else {
			fmt.Println("Config file was found but another error was produced", err)
		}
	}
	_ = viper.Unmarshal(&global.GLOBAL_CONFIG)
	fmt.Println("global.GLOBAL_CONFIG", global.GLOBAL_CONFIG)

}
